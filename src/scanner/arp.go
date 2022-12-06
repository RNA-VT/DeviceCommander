// Copyright 2012 Google, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

// arpscan implements ARP scanning of all interfaces' local networks using
// gopacket and its subpackages.  This example shows, among other things:
//   - Generating and sending packet data
//   - Reading in packet data and interpreting it
//   - Use of the 'pcap' subpackage for reading/writing
package scanner

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/rna-vt/devicecommander/src/device"
)

const (
	localHostPrefix        = 127
	arpHardwareAddressSize = 6
	protAddressSize        = 4
)

type ArpScanner struct {
	LoopDelay                    int
	NewDeviceChan                chan device.NewDeviceParams
	Stop                         chan struct{}
	PCAPPort                     int32
	BroadcastResponseWaitSeconds int
	DefaultPort                  int
	logger                       *logrus.Entry
}

func NewArpScanner(newDeviceChan chan device.NewDeviceParams, stopChan chan struct{}) ArpScanner {
	return ArpScanner{
		LoopDelay:                    viper.GetInt("ARP_SCANNER_LOOP_DELAY"),
		NewDeviceChan:                newDeviceChan,
		Stop:                         stopChan,
		PCAPPort:                     viper.GetInt32("ARP_SCANNER_PCAP_PORT"),
		BroadcastResponseWaitSeconds: viper.GetInt("ARP_SCANNER_BROADCAST_RESPONSE_WAIT_SECONDS"),
		DefaultPort:                  viper.GetInt("ARP_SCANNER_DEFAULT_DEVICE_PORT"),
		logger:                       getScannerLogger(),
	}
}

// Start is the entrypoint for the ArpScanner. It will run two concurrent
// routines for reading and writing of arp messages. The scanner can be stopped
// by closing the Stop chan.
// NewDevices will be returned via the NewDeviceChan.
func (a *ArpScanner) Start() {
	// Get a list of all interfaces.
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for _, iface := range ifaces {
		wg.Add(1)
		go func(iface net.Interface) {
			defer wg.Done()
			if err := a.scanNetworkInterface(&iface); err != nil {
				a.logger.Trace(fmt.Sprintf("interface %v: %v", iface.Name, err))
			}
		}(iface)
	}

	wg.Wait()
}

func (a *ArpScanner) validateIPAddress(ipAddress net.IPNet) error {
	if ipAddress.IP[0] == localHostPrefix {
		return errors.New("skipping localhost")
	}

	if ipAddress.Mask[0] != 0xff || ipAddress.Mask[1] != 0xff {
		return errors.New("mask means network is too large")
	}

	return nil
}

func (a *ArpScanner) hasIPv4Address(iface *net.Interface) (net.IPNet, error) {
	addrs, err := iface.Addrs()
	if err != nil {
		return net.IPNet{}, err
	}
	for _, ad := range addrs {
		if ipnet, ok := ad.(*net.IPNet); ok {
			if ip4 := ipnet.IP.To4(); ip4 != nil {
				return net.IPNet{
					IP:   ip4,
					Mask: ipnet.Mask[len(ipnet.Mask)-4:],
				}, nil
			}
		}
	}
	return net.IPNet{}, errors.New("the interface does not have any IPv4 addresses")
}

// scanNetworkInterface scans an individual interface's local network for machines using ARP
// requests/replies. It will loop forever, sending packets out regularly.  It returns an error
// if it's ever unable to write a packet.
func (a *ArpScanner) scanNetworkInterface(iface *net.Interface) error {
	ipv4Address, err := a.hasIPv4Address(iface)
	if err != nil {
		return err
	}

	if err = a.validateIPAddress(ipv4Address); err != nil {
		return err
	}

	a.logger.Info(fmt.Sprintf("Using network range %v for interface %v", ipv4Address, iface.Name))

	// Open up a PCAP handle for packet reads/writes.
	handle, err := pcap.OpenLive(iface.Name, a.PCAPPort, true, pcap.BlockForever)
	if err != nil {
		return err
	}
	defer handle.Close()

	go a.readARP(handle, iface)
	// defer close(a.Stop)
	for {
		// Write our scan packets out to the handle.
		if err := a.writeARP(handle, iface, &ipv4Address); err != nil {
			a.logger.Error(fmt.Sprintf("error writing packets on %v: %v", iface.Name, err))
			return err
		}
		// We don't know exactly how long it'll take for packets to be
		// sent back to us, but 10 seconds should be more than enough
		// time ;)
		time.Sleep(time.Duration(a.BroadcastResponseWaitSeconds) * time.Second)
	}
}

// readARP watches a handle for incoming ARP responses we might care about, and prints them.
// It will loop until 'stop' is closed.
func (a *ArpScanner) readARP(handle *pcap.Handle, iface *net.Interface) {
	src := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet)
	in := src.Packets()
	for {
		var packet gopacket.Packet
		select {
		case <-a.Stop:
			a.logger.Info("Exited readARP loop due to Stop chan closing")
			return
		case packet = <-in:
			arpLayer := packet.Layer(layers.LayerTypeARP)
			if arpLayer == nil {
				continue
			}
			arp, err := arpLayer.(*layers.ARP)
			if err {
				a.logger.Trace(fmt.Errorf("error constructing arp obj"))
			}
			if arp.Operation != layers.ARPReply || bytes.Equal([]byte(iface.HardwareAddr), arp.SourceHwAddress) {
				// This is a packet I sent.
				continue
			}

			a.logger.Debugf("ARP Packet received for [MAC: %v] [IP: %v]", net.HardwareAddr(arp.SourceHwAddress), net.IP(arp.SourceProtAddress))
			tmpMacAddress := net.HardwareAddr(arp.SourceHwAddress).String()
			ip := net.IP(arp.SourceProtAddress).String()

			newDevice := device.NewDeviceParams{
				Mac:  &tmpMacAddress,
				Host: ip,
				Port: a.DefaultPort,
			}

			// pass new device across channel
			a.NewDeviceChan <- newDevice
		}
	}
}

// writeARP writes an ARP request for each address on our local network to the
// pcap handle.
func (a *ArpScanner) writeARP(handle *pcap.Handle, iface *net.Interface, addr *net.IPNet) error {
	// Set up all the layers' fields we can.
	eth := layers.Ethernet{
		SrcMAC:       iface.HardwareAddr,
		DstMAC:       net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		EthernetType: layers.EthernetTypeARP,
	}
	arp := layers.ARP{
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     arpHardwareAddressSize,
		ProtAddressSize:   protAddressSize,
		Operation:         layers.ARPRequest,
		SourceHwAddress:   []byte(iface.HardwareAddr),
		SourceProtAddress: []byte(addr.IP),
		DstHwAddress:      []byte{0, 0, 0, 0, 0, 0},
	}
	// Set up buffer and options for serialization.
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	// Send one packet for every address.
	for _, ip := range a.ips(addr) {
		arp.DstProtAddress = []byte(ip)
		if err := gopacket.SerializeLayers(buf, opts, &eth, &arp); err != nil {
			return err
		}

		if err := handle.WritePacketData(buf.Bytes()); err != nil {
			return err
		}
	}
	return nil
}

// ips is a simple and not very good method for getting all IPv4 addresses from a
// net.IPNet.  It returns all IPs it can over the channel it sends back, closing
// the channel when done.
func (a *ArpScanner) ips(n *net.IPNet) (out []net.IP) {
	num := binary.BigEndian.Uint32([]byte(n.IP))
	mask := binary.BigEndian.Uint32([]byte(n.Mask))
	network := num & mask
	broadcast := network | ^mask
	for network++; network < broadcast; network++ {
		var buf [4]byte
		binary.BigEndian.PutUint32(buf[:], network)
		out = append(out, net.IP(buf[:]))
	}
	return
}

// Copyright 2012 Google, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

// arpscan implements ARP scanning of all interfaces' local networks using
// gopacket and its subpackages.  This example shows, among other things:
//   * Generating and sending packet data
//   * Reading in packet data and interpreting it
//   * Use of the 'pcap' subpackage for reading/writing
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
	"github.com/rna-vt/devicecommander/graph/model"
)

type ArpScanner struct {
	LoopDelay     int
	NewDeviceChan chan model.NewDevice
	Stop          chan struct{}
}

// Start is the entrypoint for the ArpScanner. It will run two concurrent
// routines for reading and writing of arp messages. The scanner can be stopped
// by closing the Stop chan.
// NewDevices will be returned via the NewDeviceChan.
func (a *ArpScanner) Start() {
	logger := getScannerLogger()
	// Get a list of all interfaces.
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for _, iface := range ifaces {
		wg.Add(1)
		// Start up a scan on each interface.
		go func(iface net.Interface) {
			defer wg.Done()
			if err := a.scan(&iface); err != nil {
				logger.Info(fmt.Sprintf("interface %v: %v", iface.Name, err))
			}
		}(iface)
	}
	// Wait for all interfaces' scans to complete.  They'll try to run
	// forever, but will stop on an error, so if we get past this Wait
	// it means all attempts to write have failed.
	wg.Wait()
}

// scan scans an individual interface's local network for machines using ARP requests/replies.
// scan loops forever, sending packets out regularly.  It returns an error if
// it's ever unable to write a packet.
func (a *ArpScanner) scan(iface *net.Interface) error {
	// We just look for IPv4 addresses, so try to find if the interface has one.
	logger := getScannerLogger()
	var addr *net.IPNet
	if addrs, err := iface.Addrs(); err != nil {
		return err
	} else {
		for _, ad := range addrs {
			if ipnet, ok := ad.(*net.IPNet); ok {
				if ip4 := ipnet.IP.To4(); ip4 != nil {
					addr = &net.IPNet{
						IP:   ip4,
						Mask: ipnet.Mask[len(ipnet.Mask)-4:],
					}
					break
				}
			}
		}
	}
	// Sanity-check that the interface has a good address.
	if addr == nil {
		return errors.New("no good IP network found")
	} else if addr.IP[0] == 127 {
		return errors.New("skipping localhost")
	} else if addr.Mask[0] != 0xff || addr.Mask[1] != 0xff {
		return errors.New("mask means network is too large")
	}
	logger.Info(fmt.Sprintf("Using network range %v for interface %v", addr, iface.Name))

	// Open up a pcap handle for packet reads/writes.
	handle, err := pcap.OpenLive(iface.Name, 65536, true, pcap.BlockForever)
	if err != nil {
		return err
	}
	defer handle.Close()

	go a.readARP(handle, iface)
	// defer close(a.Stop)
	for {
		// Write our scan packets out to the handle.
		if err := a.writeARP(handle, iface, addr); err != nil {
			logger.Error(fmt.Sprintf("error writing packets on %v: %v", iface.Name, err))
			return err
		}
		// We don't know exactly how long it'll take for packets to be
		// sent back to us, but 10 seconds should be more than enough
		// time ;)
		time.Sleep(60 * time.Second)
	}
}

// readARP watches a handle for incoming ARP responses we might care about, and prints them.
// readARP loops until 'stop' is closed.
func (a *ArpScanner) readARP(handle *pcap.Handle, iface *net.Interface) {
	logger := getScannerLogger()
	src := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet)
	in := src.Packets()
	for {
		var packet gopacket.Packet
		select {
		case <-a.Stop:
			logger.Info("Exited readARP loop due to Stop chan closing")
			return
		case packet = <-in:
			arpLayer := packet.Layer(layers.LayerTypeARP)
			if arpLayer == nil {
				continue
			}
			arp, err := arpLayer.(*layers.ARP)
			if err {
				logger.Trace(fmt.Errorf("error constructing arp obj"))
			}
			if arp.Operation != layers.ARPReply || bytes.Equal([]byte(iface.HardwareAddr), arp.SourceHwAddress) {
				// This is a packet I sent.
				continue
			}
			// Note:  we might get some packets here that aren't responses to ones we've sent,
			// if for example someone else sends US an ARP request.  Doesn't much matter, though...
			// all information is good information :)
			logger.Debug(fmt.Sprintf("IP %v is at %v", net.IP(arp.SourceProtAddress), net.HardwareAddr(arp.SourceHwAddress)))
			tmpMacAddress := net.HardwareAddr(arp.SourceHwAddress).String()

			newDevice := model.NewDevice{
				Mac:  &tmpMacAddress,
				Host: net.IP(arp.SourceProtAddress).String(),
				Port: 420,
			}

			// passs new device across channel
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
		HwAddressSize:     6,
		ProtAddressSize:   4,
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
		gopacket.SerializeLayers(buf, opts, &eth, &arp)
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

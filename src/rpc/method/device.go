package method

import "github.com/rna-vt/devicecommander/src/rpc"

type GetDevicePayload struct {
	rpc.JSONRPC
	// example: getDevice
	Method string `json:"method"`
	// example: 705e4dcb-3ecd-24f3-3a35-3e926e4bded5
	Params string `json:"params"`
}

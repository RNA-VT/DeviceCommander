package rpc

type JSONRPC struct {
	// example: 1.0
	Version string `json:"jsonrpc"`

	// example: 1
	ID int `json:"id"`

	// example: doMethod
	Method string `json:"method"`
}
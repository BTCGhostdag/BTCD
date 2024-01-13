package rpchandlers

import (
	"github.com/BTCGhostdag/BTCD/app/appmessage"
	"github.com/BTCGhostdag/BTCD/app/rpc/rpccontext"
	"github.com/BTCGhostdag/BTCD/infrastructure/network/netadapter/router"
)

// HandleGetHeaders handles the respectively named RPC command
func HandleGetHeaders(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetHeadersResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}

package rpchandlers

import (
	"github.com/BTCGhostdag/BTCD/app/appmessage"
	"github.com/BTCGhostdag/BTCD/app/rpc/rpccontext"
	"github.com/BTCGhostdag/BTCD/domain/consensus/utils/constants"
	"github.com/BTCGhostdag/BTCD/infrastructure/network/netadapter/router"
)

// HandleGetCoinSupply handles the respectively named RPC command
func HandleGetCoinSupply(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	if !context.Config.UTXOIndex {
		errorMessage := &appmessage.GetCoinSupplyResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Method unavailable when BTCD is run without --utxoindex")
		return errorMessage, nil
	}

	circulatingSompiSupply, err := context.UTXOIndex.GetCirculatingSompiSupply()
	if err != nil {
		return nil, err
	}

	response := appmessage.NewGetCoinSupplyResponseMessage(
		constants.MaxSompi,
		circulatingSompiSupply,
	)

	return response, nil
}

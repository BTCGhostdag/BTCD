package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/BTCGhostdag/BTCD/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.BTCDdMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.BTCDdMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.BTCDdMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.BTCDdMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.BTCDdMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.BTCDdMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.BTCDdMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.BTCDdMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.BTCDdMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.BTCDdMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.BTCDdMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.BTCDdMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.BTCDdMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.BTCDdMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.BTCDdMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.BTCDdMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.BTCDdMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.BTCDdMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.BTCDdMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.BTCDdMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.BTCDdMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.BTCDdMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.BTCDdMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.BTCDdMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.BTCDdMessage_BanRequest{}),
	reflect.TypeOf(protowire.BTCDdMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}

package flowcontext

import (
	"github.com/BTCGhostdag/BTCD/infrastructure/network/addressmanager"
)

// AddressManager returns the address manager associated to the flow context.
func (f *FlowContext) AddressManager() *addressmanager.AddressManager {
	return f.addressManager
}

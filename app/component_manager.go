package app

import (
	"fmt"
	"sync/atomic"

	"github.com/BTCGhostdag/BTCD/domain/consensus/model/externalapi"

	"github.com/BTCGhostdag/BTCD/domain/miningmanager/mempool"

	"github.com/BTCGhostdag/BTCD/app/protocol"
	"github.com/BTCGhostdag/BTCD/app/rpc"
	"github.com/BTCGhostdag/BTCD/domain"
	"github.com/BTCGhostdag/BTCD/domain/consensus"
	"github.com/BTCGhostdag/BTCD/domain/utxoindex"
	"github.com/BTCGhostdag/BTCD/infrastructure/config"
	infrastructuredatabase "github.com/BTCGhostdag/BTCD/infrastructure/db/database"
	"github.com/BTCGhostdag/BTCD/infrastructure/network/addressmanager"
	"github.com/BTCGhostdag/BTCD/infrastructure/network/connmanager"
	"github.com/BTCGhostdag/BTCD/infrastructure/network/netadapter"
	"github.com/BTCGhostdag/BTCD/infrastructure/network/netadapter/id"
	"github.com/BTCGhostdag/BTCD/util/panics"
)

// ComponentManager is a wrapper for all the BTCD services
type ComponentManager struct {
	cfg               *config.Config
	addressManager    *addressmanager.AddressManager
	protocolManager   *protocol.Manager
	rpcManager        *rpc.Manager
	connectionManager *connmanager.ConnectionManager
	netAdapter        *netadapter.NetAdapter

	started, shutdown int32
}

// Start launches all the BTCD services.
func (a *ComponentManager) Start() {
	// Already started?
	if atomic.AddInt32(&a.started, 1) != 1 {
		return
	}

	log.Trace("Starting BTCD")

	err := a.netAdapter.Start()
	if err != nil {
		panics.Exit(log, fmt.Sprintf("Error starting the net adapter: %+v", err))
	}

	a.connectionManager.Start()
}

// Stop gracefully shuts down all the BTCD services.
func (a *ComponentManager) Stop() {
	// Make sure this only happens once.
	if atomic.AddInt32(&a.shutdown, 1) != 1 {
		log.Infof("BTCDd is already in the process of shutting down")
		return
	}

	log.Warnf("BTCDd shutting down")

	a.connectionManager.Stop()

	err := a.netAdapter.Stop()
	if err != nil {
		log.Errorf("Error stopping the net adapter: %+v", err)
	}

	a.protocolManager.Close()
	close(a.protocolManager.Context().Domain().ConsensusEventsChannel())

	return
}

// NewComponentManager returns a new ComponentManager instance.
// Use Start() to begin all services within this ComponentManager
func NewComponentManager(cfg *config.Config, db infrastructuredatabase.Database, interrupt chan<- struct{}) (
	*ComponentManager, error) {

	consensusConfig := consensus.Config{
		Params:                          *cfg.ActiveNetParams,
		IsArchival:                      cfg.IsArchivalNode,
		EnableSanityCheckPruningUTXOSet: cfg.EnableSanityCheckPruningUTXOSet,
	}
	mempoolConfig := mempool.DefaultConfig(&consensusConfig.Params)
	mempoolConfig.MaximumOrphanTransactionCount = cfg.MaxOrphanTxs
	mempoolConfig.MinimumRelayTransactionFee = cfg.MinRelayTxFee

	domain, err := domain.New(&consensusConfig, mempoolConfig, db)
	if err != nil {
		return nil, err
	}

	netAdapter, err := netadapter.NewNetAdapter(cfg)
	if err != nil {
		return nil, err
	}

	addressManager, err := addressmanager.New(addressmanager.NewConfig(cfg), db)
	if err != nil {
		return nil, err
	}

	var utxoIndex *utxoindex.UTXOIndex
	if cfg.UTXOIndex {
		utxoIndex, err = utxoindex.New(domain, db)
		if err != nil {
			return nil, err
		}

		log.Infof("UTXO index started")
	}

	connectionManager, err := connmanager.New(cfg, netAdapter, addressManager)
	if err != nil {
		return nil, err
	}
	protocolManager, err := protocol.NewManager(cfg, domain, netAdapter, addressManager, connectionManager)
	if err != nil {
		return nil, err
	}
	rpcManager := setupRPC(cfg, domain, netAdapter, protocolManager, connectionManager, addressManager, utxoIndex, domain.ConsensusEventsChannel(), interrupt)

	return &ComponentManager{
		cfg:               cfg,
		protocolManager:   protocolManager,
		rpcManager:        rpcManager,
		connectionManager: connectionManager,
		netAdapter:        netAdapter,
		addressManager:    addressManager,
	}, nil

}

func setupRPC(
	cfg *config.Config,
	domain domain.Domain,
	netAdapter *netadapter.NetAdapter,
	protocolManager *protocol.Manager,
	connectionManager *connmanager.ConnectionManager,
	addressManager *addressmanager.AddressManager,
	utxoIndex *utxoindex.UTXOIndex,
	consensusEventsChan chan externalapi.ConsensusEvent,
	shutDownChan chan<- struct{},
) *rpc.Manager {

	rpcManager := rpc.NewManager(
		cfg,
		domain,
		netAdapter,
		protocolManager,
		connectionManager,
		addressManager,
		utxoIndex,
		consensusEventsChan,
		shutDownChan,
	)
	protocolManager.SetOnNewBlockTemplateHandler(rpcManager.NotifyNewBlockTemplate)
	protocolManager.SetOnPruningPointUTXOSetOverrideHandler(rpcManager.NotifyPruningPointUTXOSetOverride)

	return rpcManager
}

// P2PNodeID returns the network ID associated with this ComponentManager
func (a *ComponentManager) P2PNodeID() *id.ID {
	return a.netAdapter.ID()
}

// AddressManager returns the AddressManager associated with this ComponentManager
func (a *ComponentManager) AddressManager() *addressmanager.AddressManager {
	return a.addressManager
}

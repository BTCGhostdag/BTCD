package model

import "github.com/BTCGhostdag/BTCD/domain/consensus/model/externalapi"

// BlockStatusStore represents a store of BlockStatuses
type BlockStatusStore interface {
	Store
	Stage(stagingArea *StagingArea, blockHash *externalapi.DomainHash, blockStatus externalapi.BlockStatus)
	IsStaged(stagingArea *StagingArea) bool
	Get(dbContext DBReader, stagingArea *StagingArea, blockHash *externalapi.DomainHash) (externalapi.BlockStatus, error)
	Exists(dbContext DBReader, stagingArea *StagingArea, blockHash *externalapi.DomainHash) (bool, error)
}

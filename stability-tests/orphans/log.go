package main

import (
	"github.com/BTCGhostdag/BTCD/infrastructure/logger"
	"github.com/BTCGhostdag/BTCD/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("ORPH")
	spawn      = panics.GoroutineWrapperFunc(log)
)

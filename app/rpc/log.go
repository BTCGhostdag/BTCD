package rpc

import (
	"github.com/BTCGhostdag/BTCD/infrastructure/logger"
	"github.com/BTCGhostdag/BTCD/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)

package rpcclient

import (
	"github.com/BTCGhostdag/BTCD/infrastructure/logger"
	"github.com/BTCGhostdag/BTCD/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)

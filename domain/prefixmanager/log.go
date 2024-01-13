package prefixmanager

import (
	"github.com/BTCGhostdag/BTCD/infrastructure/logger"
	"github.com/BTCGhostdag/BTCD/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)

package standalone

import (
	"github.com/BTCGhostdag/BTCD/infrastructure/logger"
	"github.com/BTCGhostdag/BTCD/util/panics"
)

var log = logger.RegisterSubSystem("NTAR")
var spawn = panics.GoroutineWrapperFunc(log)

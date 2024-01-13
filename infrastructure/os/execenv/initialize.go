package execenv

import (
	"fmt"
	"os"
	"runtime"

	"github.com/BTCGhostdag/BTCD/infrastructure/os/limits"
)

// Initialize initializes the execution environment required to run BTCD
func Initialize(desiredLimits *limits.DesiredLimits) {
	// Use all processor cores.
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Up some limits.
	if err := limits.SetLimits(desiredLimits); err != nil {
		fmt.Fprintf(os.Stderr, "failed to set limits: %s\n", err)
		os.Exit(1)
	}

}

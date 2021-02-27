package cyclemod2

import (
	"fmt"

	"changkun.de/x/cyclemod"
)

// Call prints the depend version of cyclemod
func Call() {
	fmt.Printf("cyclemod2@%s depends cyclemod@%s\n",
		Version, cyclemod.Version)
}

// Version is the version of the module
var Version = "v2.1.0"

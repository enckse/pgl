// Handles testing the fmt.Die* calls
package main

import (
	"flag"

	fmt "github.com/enckse/pgl/exit"
)

func main() {
	mode := flag.String("mode", "", "testing mode")
	flag.Parse()
	switch *mode {
	case "dief":
		fmt.Dief("dief: %s %d", "test", 1)
	case "die":
		fmt.Die("die")
	case "dieNil":
		fmt.Die(nil)
	case "dieExit":
		fmt.DieAndExit(2, "die-and-exit")
	case "dieExitNil":
		fmt.DieAndExit(3, nil)
	case "diefExit":
		fmt.DieAndExitf(4, "dief-and-exit: %t", true)
	default:
		fmt.Dief("unknown mode")
	}
}

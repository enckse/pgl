// Handles testing the fmt.Die* calls
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/enckse/pgl/os/stdin"
)

func main() {
	mode := flag.String("mode", "", "testing mode")
	flag.Parse()
	var b []byte
	var err error
	switch *mode {
	case "one":
		b, err = stdin.ReadLine()
	case "multiple":
		b, err = stdin.ReadAll()
	default:
		os.Exit(1)
	}
	fmt.Printf("data: %s\n", string(b))
	fmt.Printf("error: %v\n", err)
}

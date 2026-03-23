// emmp is the Emmp developer CLI for Cedar policy tooling.
//
// Commands:
//
//	emmp policy validate <file.cedar>
//	emmp policy test [--tier unit|regression|contract]
//	emmp policy lint <dir>
//	emmp policy scan <namespace>
//	emmp policy analyze <namespace> --query what-can/who-can
package main

import (
	"fmt"
	"os"
)

var version = "dev"

func main() {
	if err := rootCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

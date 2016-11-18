// echo.go Print cmdline args with flags.

package main

import (
	"flag"
	"fmt"
	"strings"
)

// Create flags for use in cmdline args

// Flag of type Bool
var n = flag.Bool("n", false, "omit trailing spaces")

// Flag of type String
var sep = flag.String("s", " ", "seperator")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))

	// sep, n are pointers to Flag
	if !*n {
		fmt.Println()
	}
}

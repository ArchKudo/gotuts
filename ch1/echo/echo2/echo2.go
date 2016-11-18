// 1.2 echo2.go Print commandline arguments using range.

package main

import (
	"fmt"
	"os"
)

func main() {

	// Declare both together
	// Shorthand can be used only inside functions and not at package-level
	s, sep := "", ""

	// _ for blank identifier
	// range produces a pair of values: index, value at that index
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// boiling.go Print the boiling temperature of water.

package main

import "fmt"

// A package-level declaration available throughout package source files
const boilingF = 212.0

func main() {
	// A local variable declaration
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("The boiling temperature of water is %g°F or %g°C\n", f, c)
}

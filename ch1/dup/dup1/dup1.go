// dup1.go Prints duplicate strings and their count.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// map[key]value is unordered
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)

	// Scan() advances scanner to new token
	for input.Scan() {

		// Gets string, Creates key in counts, Increases index for similar string
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d: %s\n", n, line)
		}
	}
}

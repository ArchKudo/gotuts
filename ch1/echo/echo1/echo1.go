// 1.1 echo1.go Print commandline arguments using for loop.

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	// for init; condition; post {..} (for loop)
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	// for condition {..} (while loop)
	i := 1
	for i < len(os.Args) {
		fmt.Println(os.Args[i])
		i++
	}

	// for {..} (infinite loop)
	i = 1
	for {
		if i < len(os.Args) {
			fmt.Println(os.Args[i])
			i++
		} else {
			break
		}
	}
}

// exercises.go Covers Exercises 1.1, 1.2, 1.3

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// 1.1, 1.2
	for i, arg := range os.Args[0:] {
		fmt.Printf("%v: %v\n", i, arg)
	}

	// 1.3 For loop
	var s, sep string
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	elapsed := time.Since(start).Seconds()
	fmt.Printf("Elapsed time: %fs\n", elapsed)

	// 1.3 String.join
	start = time.Now()
	s = strings.Join(os.Args[1:], " ")
	fmt.Println(s)
	elapsed = time.Since(start).Seconds()
	fmt.Printf("Elapsed time: %fs\n", elapsed)
}

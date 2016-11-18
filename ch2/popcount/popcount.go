// popcount.go Checks whether an uint64 is initialized.

package main

import (
	"fmt"
)

var pc [256]byte

// initialize pc
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// popCount returns the population count (number of set bits) of x
func popCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	var x uint64 = 1
	fmt.Println(pc)
	y := popCount(x)
	fmt.Println(y)
}

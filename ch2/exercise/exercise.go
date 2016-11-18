// popcount.go Checks whether an uint64 is initialized.

package main

import (
	"fmt"
	"time"
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

func popCount2(x uint64) int {
	var a int
	var i uint
	for i = 0; i <= 8; i++ {
		a += int(pc[byte(x>>(i*8))])
	}
	return a
}

func main() {
	var x uint64 = 1
	fmt.Println(pc)
	start := time.Now()
	y := popCount(x)
	sec := time.Since(start).Seconds()
	fmt.Printf("Initialized = %d, time = %.12f\n", y, sec)
	start = time.Now()
	y = popCount2(x)
	sec = time.Since(start).Seconds()
	fmt.Printf("Initialized = %d, time = %.12f\n", y, sec)
}

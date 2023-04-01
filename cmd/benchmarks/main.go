package main

import (
	"fmt"
	"math/big"
)

// FibInt will compute fibinochi sequence using machine native integers
func FibInt(n int) int {
	if n < 2 {
		return n
	}
	a, b := 0, 1
	for n--; n > 0; n-- {
		a += b
		a, b = b, a
	}
	return b
}

// FibBig will compute fibinochi sequence using Go math/big package
func FibBig(n int) *big.Int {
	if n < 2 {
		return big.NewInt(int64(n))
	}
	a, b := big.NewInt(0), big.NewInt(1)
	for n--; n > 0; n-- {
		a.Add(a, b)
		a, b = b, a
	}
	return b
}

func main() {
	fmt.Printf("FibInt: %v", FibInt(30))
	fmt.Printf("FibBig: %v", FibBig(30))
}

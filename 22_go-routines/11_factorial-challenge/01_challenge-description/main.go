// Package main provides challenge description
package main

import (
	"fmt"
)

func main() {
	f := factorial(4)
	fmt.Println(f)
}

func factorial(n int) int {
	fact := 1
	for i := n; i > 0; i-- {
		fact *= i
	}
	return fact
}

/*
CHALLENGE 1:
-- Use go routine and channels to calculate factorial

CHALLENGE 2:
--
*/

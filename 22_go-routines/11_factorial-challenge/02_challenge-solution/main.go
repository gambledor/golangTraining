// Package main provides factorial challenge solution.
package main

import (
	"fmt"
)

func main() {
	cin := factorial(4)
	for n := range cin {
		fmt.Println(n)
	}
}

func factorial(n int) <-chan int {
	cout := make(chan int)
	go func() {
		f := 1
		for i := n; i > 0; i-- {
			f *= i
		}
		cout <- f
		close(cout)
	}()
	return cout
}

// Package main provides exercise 6
package main

import "fmt"

func main() {
	var sum int

	for i := 0; i < 1000; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}
	fmt.Printf("The sum of all multiples of 3 or 5 below 1000 is: %d\n", sum)
}

// Package main provides exercise 1
package main

import "fmt"

func main() {
	var smaller, larger int

	fmt.Println("Enter a small number:")
	fmt.Scan(&smaller)
	fmt.Println("Enter a large  number:")
	fmt.Scan(&larger)
	fmt.Printf("Remainder %d\n", larger%smaller)
}

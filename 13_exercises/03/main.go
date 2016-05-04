// Package main provides exercise 1
package main

import "fmt"

func main() {
	var username string

	fmt.Println("Enter your name:")
	fmt.Scan(&username)
	fmt.Printf("Hello %s\n", username)
}

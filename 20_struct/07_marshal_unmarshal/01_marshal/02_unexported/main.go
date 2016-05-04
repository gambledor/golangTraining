// Package main provides ...
package main

import (
	"encoding/json"
	"fmt"
)

// Person struct
type Person struct {
	first, last string
	age         int
}

func main() {
	p1 := Person{"James", "Bond", 20}
	fmt.Println(p1)
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}

// Package main provides ...
package main

import (
	"encoding/json"
	"fmt"
)

// Person struct
type Person struct {
	First, Last      string
	Age, notExported int
}

func main() {
	p1 := Person{"James", "Bond", 20, 007}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}

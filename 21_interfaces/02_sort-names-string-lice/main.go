// Package main provides ...
package main

import (
	"fmt"
	"sort"
)

// func (s []string) Len() int           { len(s) }
// func (s []string) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
// func (s []string) Less(i, j int) bool { s[i] < s[j] }
//
func main() {
	var s = []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	// sort.StringSlice(s).Sort()
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)
}

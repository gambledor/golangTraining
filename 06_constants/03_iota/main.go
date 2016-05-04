package main

import "fmt"

const (
    _ = iota //0
    Kb = 1 << (iota * 10) // 1 << (1 * 10)
    Mb = 1 << (iota * 10) // 1 << (2 * 10)
    Gb = 1 << (iota * 10) // 1 << (3 * 10)
    Tb = 1 << (iota * 10) // 1 << (4 * 10)
)

func main()  {
    fmt.Println("binary\t\tdecimal")
    fmt.Printf("%b\t", Kb)
    fmt.Printf("%d\n", Kb)
    fmt.Printf("%b\t", Mb)
    fmt.Printf("%d\n", Mb)
    fmt.Printf("%b\t", Gb)
    fmt.Printf("%d\n", Gb)
    fmt.Printf("%b\t", Tb)
    fmt.Printf("%d\n", Tb)
}


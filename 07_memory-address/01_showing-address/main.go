package main

import "fmt"

func main()  {
    var a = 43

    fmt.Println("a - ", a)
    fmt.Println("a's memory address - ", &a)
    fmt.Printf("%d\n", &a)
}


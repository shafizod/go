package main

import "fmt"

func main() {
    var a, b int
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("B = ")
    fmt.Scan(&b)
    result := a >= 0 || b < -2
    fmt.Printf("result = %t\n", result)
}
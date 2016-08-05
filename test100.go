// test project main.go
package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
    fmt.Printf("%d ", i)
    if i % 3 == 0 {
        fmt.Printf("Fizz")
    }
    if i % 5 == 0 {
        fmt.Printf("Buzz")
    }
    fmt.Println()
fmt.Scanln()
}
}
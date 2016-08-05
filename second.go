// test project main.go
package main

import (
	"fmt"
)

func main() {
	for N := 1; N <= 100; N++ {
	if N % 3 == 0 {
	fmt.Println(N)
	}
	}
fmt.Scanln()
}

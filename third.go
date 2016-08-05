// test project main.go
package main

import (
	"fmt"
)

func main() {
	for N := 1; N <= 100; N++ {
		fmt.Println()
		if N%3 != 0 && N%5 != 0 {
			fmt.Print(N)
			continue
		}
		if N%3 == 0 {
			fmt.Print("Fizz")
		}
		if N%5 == 0 {
			fmt.Print("Buzz")
		}
	}
	fmt.Scanln()
}

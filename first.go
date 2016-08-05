// test project main.go
package main

import (
	"fmt"
)

func main() {
	N := 0
	for {
	fmt.Println(N)
	N = N + 1
	if N == 10 {
	break
	}	
}
fmt.Scanln()
}

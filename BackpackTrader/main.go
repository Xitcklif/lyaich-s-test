package main

import (
	"fmt"
	"os"
)

func main() {
	excel_file := os.Args[1]
	fmt.Printf("Processing excel file: %s\n", excel_file)
}

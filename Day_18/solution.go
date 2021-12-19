package main

import (
	"fmt"
	"os"
)

func main() {
    data, err := LoadData("input.txt")
    if err != nil {
        fmt.Printf("Error loading input")
        os.Exit(1)
    }
    d := AddNumbers(data)
    fmt.Printf("Solution: %d", magnitude(d))
}

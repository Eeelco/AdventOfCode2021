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
    fmt.Printf("Part 1 solution: %d\n\n", magnitude(d))
    l := LargestMagnitude(data)
    fmt.Printf("Part 2 solution: %d\n\n", l)
}

func LargestMagnitude(numbers []string) int {
    largest := -1
    for i := 0; i < len(numbers); i++ {
        for j := 0; j < len(numbers); j++ {
            if i != j {
                sum := add(MakeTree(numbers[i]), MakeTree(numbers[j]))
                reduce(sum)
                m := magnitude(sum)
                if m > largest {
                    largest = m
                }
            }
        }
    }
    return largest
}

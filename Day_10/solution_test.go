package day10

import (
	"fmt"
	"testing"
)

func Test_Day6_Part1(t *testing.T) {
    data, err := LoadData("test_data.txt")
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    Corr := CorruptedErrors(data)
    if Corr != 26397 {
        t.Fatalf("Test data error. Expected 26397, got %d\n", Corr)
    }

    data, err = LoadData("input.txt")
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    Corr = CorruptedErrors(data)
    fmt.Printf("Part 1 solution:\n\n%d\n\n",Corr)
}

func Test_Day6_Part2(t *testing.T) {
    data, err := LoadData("test_data.txt")
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    score := FixIncomplete(data)
    if score != 288957 {
        t.Fatalf("Test data error. Expected 288957, got %d\n", score)
    }

    data, err = LoadData("input.txt")
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    score = FixIncomplete(data)
    fmt.Printf("Part 2 solution:\n\n%d\n\n",score)
}

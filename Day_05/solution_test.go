package day05

import (
	"fmt"
	"testing"
)


func Test_Day5_Part1(t *testing.T) {
    lines, err := LoadData("test_data.txt")
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    intersects := GetResult(lines, true)
    if intersects != 5 {
        t.Fatalf("Test data error. Expected 5, got %d\n", intersects)
    }

    lines, err = LoadData("input.txt")
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    intersects = GetResult(lines, true)
    fmt.Printf("Part 1 solution:\nOverlaps = %d\n\n", intersects)
}

func Test_Day5_Part2(t *testing.T) {
    lines, err := LoadData("test_data.txt")
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    intersects := GetResult(lines, false)
    if intersects != 12 {
        t.Fatalf("Test data error. Expected 12, got %d\n", intersects)
    }

    lines, err = LoadData("input.txt")
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    intersects = GetResult(lines, false)
    fmt.Printf("Part 2 solution:\nOverlaps = %d\n\n", intersects)
}

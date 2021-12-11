package day11

import (
	"fmt"
	"testing"
)

func Test_Day6_Part1(t *testing.T) {
    flash_count, err := Simulate("test_data.txt", 100)
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    if flash_count != 1656{
        t.Fatalf("Test data error. Expected 1656, got %d\n", flash_count)
    }

    flash_count, err = Simulate("input.txt", 100)
    if err != nil {
        t.Fatalf("Error opening input file: %v", err)
    }
    fmt.Printf("Part 1 solution:\n\n%d\n\n",flash_count)
}

func Test_Day6_Part2(t *testing.T) {
    flash_count, err := Synchronization("test_data.txt")
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    if flash_count != 195{
        t.Fatalf("Test data error. Expected 195, got %d\n", flash_count)
    }

    flash_count, err = Synchronization("input.txt")
    if err != nil {
        t.Fatalf("Error opening input file: %v", err)
    }
    fmt.Printf("Part 2 solution:\n\n%d\n\n",flash_count)
}

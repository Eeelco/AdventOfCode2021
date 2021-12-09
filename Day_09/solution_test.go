package day09

import (
	"fmt"
	"testing"
)

func Test_Day6_Part1(t *testing.T) {
    data, err := LoadData("test_data.txt")
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    danger := count_local_minima(data)
    if danger != 15 {
        t.Fatalf("Test data error. Expected 15, got %d\n", danger)
    }

    data, err = LoadData("input.txt")
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    danger = count_local_minima(data)
    fmt.Printf("Part 1 solution:\n\n%d\n\n",danger)
}


func Test_Day6_Part2(t *testing.T) {
    data, err := LoadData("test_data.txt")
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    basins := find_basins(data)
    if basins != 1134 {
        t.Fatalf("Test data error. Expected 1134, got %d\n", basins)
    }

    data, err = LoadData("input.txt")
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    basins = find_basins(data)
    fmt.Printf("Part 1 solution:\n\n%d\n\n",basins)
}

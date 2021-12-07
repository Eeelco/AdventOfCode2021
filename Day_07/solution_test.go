package day07

import (
	"fmt"
	"testing"
)

func Test_Day6_Part1(t *testing.T) {
	data, err := ReadFile("test_data.txt")
	if err != nil {
		t.Fatalf("Error opening test file: %v", err)
	}
	dist_sum := CalcDistances(data)
	if dist_sum != 37 {
		t.Fatalf("Test data error. Expected 37, got %d\n", dist_sum)
	}

	data, err = ReadFile("input")
	if err != nil {
		t.Fatalf("Error opening test file: %v", err)
	}
	dist_sum = CalcDistances(data)
	fmt.Printf("Part 1 solution:\nFuel = %d\n\n", dist_sum)
}

func Test_Day6_Part2(t *testing.T) {
	data, err := ReadFile("test_data.txt")
	if err != nil {
		t.Fatalf("Error opening test file: %v", err)
	}
	dist_sum := GetBestMedian(data)
	if dist_sum != 168 {
		t.Fatalf("Test data error. Expected 168, got %d\n", dist_sum)
	}

	data, err = ReadFile("input")
	if err != nil {
		t.Fatalf("Error opening test file: %v", err)
	}
	dist_sum = GetBestMedian(data)
	fmt.Printf("Part 2 solution:\nFuel = %d\n\n", dist_sum)
}

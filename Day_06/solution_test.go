package day06

import (
	"fmt"
	"testing"
)

func Test_Day5_Part1(t *testing.T) {
	initial_data, err := ReadFile("test_data.txt")
	if err != nil {
		t.Fatalf("Error opening test file: %v", err)
	}
	pop := RunSimulations(initial_data, 80)
	if pop != 5934 {
		t.Fatalf("Test data error. Expected 5934, got %d\n", pop)
	}

	initial_data, err = ReadFile("input.txt")
	if err != nil {
		t.Fatalf("Error opening test file: %v", err)
	}
	pop = RunSimulations(initial_data, 80)
	fmt.Printf("Part 1 solution:\nPopulation = %d\n\n", pop)
}

func Test_Day5_Part2(t *testing.T) {
	initial_data, err := ReadFile("test_data.txt")
	if err != nil {
		t.Fatalf("Error opening test file: %v", err)
	}
	pop := RunSimulations(initial_data, 256)
	if pop != 26984457539 {
		t.Fatalf("Test data error. Expected 26984457539, got %d\n", pop)
	}

	initial_data, err = ReadFile("input.txt")
	if err != nil {
		t.Fatalf("Error opening test file: %v", err)
	}
	pop = RunSimulations(initial_data, 256)
	fmt.Printf("Part 2 solution:\nPopulation = %d\n\n", pop)
}

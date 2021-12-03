package day03

import (
	"fmt"
	"testing"
)

func Test_Day3_Part1(t *testing.T) {
	test_data, err := ReadFile("test_data.txt")
	if err != nil {
		t.Fatalf("Error opening test file: %v", err)
	}
	epsilon, gamma := CalcRates(test_data)
	if epsilon != 9 || gamma != 22 {
		t.Fatalf("Expected eps:%d gamma:%d\nGot: eps:%d gamma%d", epsilon, gamma, epsilon, gamma)
	}

	real_data, err := ReadFile("input.txt")
	if err != nil {
		t.Fatalf("Error opening test file: %v", err)
	}
	epsilon, gamma = CalcRates(real_data)
	fmt.Printf("Solution part 1:\neps:%d gamma:%d\nProduct: %d\n", epsilon, gamma, epsilon*gamma)
}

func Test_Day3_Part2(t *testing.T) {
	test_data, err := ReadFile("test_data.txt")
	if err != nil {
		t.Fatalf("Error opening test file: %v", err)
	}
	co, oxygen := GeneratorRatings(test_data)
	if co != 10 || oxygen != 23 {
		t.Fatalf("Expected co: 10 oxygen: 23\nGot: co:%d oxygen:%d", co, oxygen)
	}

	real_data, err := ReadFile("input.txt")
	if err != nil {
		t.Fatalf("Error opening test file: %v", err)
	}
	co, oxygen = GeneratorRatings(real_data)
	fmt.Printf("\n\nSolution part 2:\nco:%d oxygen:%d\nProduct: %d\n\n", co, oxygen, co*oxygen)
}

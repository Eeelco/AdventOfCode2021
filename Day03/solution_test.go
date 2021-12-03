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

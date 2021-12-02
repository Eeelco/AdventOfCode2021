package day02

import (
	"fmt"
	"testing"
)

func Test_Day2_Part1(t *testing.T) {
	test_x, test_y := 15, 10
	test_data, err := ReadFile("test_data.txt")
	if err != nil {
		t.Fatalf("Error opening test file: %v", err)
	}
	found_x, found_y := FindEndpoint(&test_data, 0, 0)
	if found_x != test_x || found_y != test_y {
		t.Fatalf("Expected x:%d y%d\nGot: x:%d y%d", test_x, test_y, found_x, found_y)
	}
	real_data, err := ReadFile("input.txt")
	if err != nil {
		t.Fatalf("Error opening data file: %v", err)
	}
	found_x, found_y = FindEndpoint(&real_data, 0, 0)

	fmt.Printf("Solution part 1:\nx:%d y:%d\nProduct: %d\n", found_x, found_y, found_x*found_y)
}

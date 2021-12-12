package day12

import (
	"fmt"
	"testing"
)

func Test_Day6(t *testing.T) {
    path_count, path_count_2, err := GetAllPaths("test_data.txt")
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    if path_count != 10{
        t.Fatalf("Test data error. Expected 10, got %d\n", path_count)
    }
    if path_count_2 != 36{
        t.Fatalf("Test data error. Expected 36, got %d\n", path_count_2)
    }
    path_count, path_count_2, err = GetAllPaths("test_data_2.txt")
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    if path_count != 19{
        t.Fatalf("Test data error. Expected 19, got %d\n", path_count)
    }
    if path_count_2 != 103{
        t.Fatalf("Test data error. Expected 103, got %d\n", path_count_2)
    }
    path_count, path_count_2, err = GetAllPaths("test_data_3.txt")
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    if path_count != 226{
        t.Fatalf("Test data error. Expected 226, got %d\n", path_count)
    }
    if path_count_2 != 3509{
        t.Fatalf("Test data error. Expected 3509, got %d\n", path_count_2)
    }

    path_count, path_count_2, err = GetAllPaths("input.txt")
    if err != nil {
        t.Fatalf("Error opening input file: %v", err)
    }
    fmt.Printf("Part 1 solution:\n\n%d\n\n",path_count)
    fmt.Printf("Part 2 solution:\n\n%d\n\n",path_count_2)
}

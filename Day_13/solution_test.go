package day13

import (
	"fmt"
	"testing"
)


func Test_Day6(t *testing.T) {
    points, folds, err := LoadData("test_data.txt")
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    after_one := PerformFold(points, folds[0])
    if len(after_one) != 17 {
        t.Fatalf("Test error. Expected 17, got %d\n", len(after_one))
    }

    points, folds, err = LoadData("input.txt")
    if err != nil {
        t.Fatalf("Error opening input file: %v", err)
    }
    after_one = PerformFold(points, folds[0])
    fmt.Printf("Part 1 solution:\n\n%d\n\n",len(after_one))
}

func Test_Day13_part2(t *testing.T) {
    points, folds, err := LoadData("test_data.txt")
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    fmt.Printf("Part 2 test:\n")
    GetSolution(points, folds)
    fmt.Printf("Real data:\n")

    points, folds, err = LoadData("input.txt")
    GetSolution(points, folds)
}

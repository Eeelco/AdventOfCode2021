package day08

import (
	"fmt"
	"testing"
)


func Test_Day6_Part1(t *testing.T) {
    data, err := LoadData("test_data.txt")
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    count_1478 := Count_1478(data)
    if count_1478 != 26 {
        t.Fatalf("Test data error. Expected 26, got %d\n", count_1478)
    }

    data, err = LoadData("input.txt")
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    count_1478 = Count_1478(data)
    fmt.Printf("Part 1 solution:\n\n%d\n\n",count_1478)
}


func Test_Day6_Part2(t *testing.T) {
    data, err := LoadData("test_data.txt")
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    digits, output := SplitLines(data)
    result := 0
    for i,f := range digits {
        mapping := DeduceDigits(f)
        val := FindNumber(output[i], mapping)
        result += val
    }
    if result != 61229 {
        t.Fatalf("Test data error. Expected 61229, got %d\n", result)
    }
    
    data, err = LoadData("input.txt")
    if err != nil {
        t.Fatalf("Error opening data file: %v", err)
    }
    digits, output = SplitLines(data)
    result = 0
    for i,f := range digits {
        mapping := DeduceDigits(f)
        val := FindNumber(output[i], mapping)
        result += val
    }
    fmt.Printf("Part 2 solution:\n\n%d\n\n", result)
}

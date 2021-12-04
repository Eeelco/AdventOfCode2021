package day04

import (
	"fmt"
	"testing"
)

func Test_Day4_Part1(t *testing.T) {
    moves, boards, err := LoadData("test_data.txt")
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    final_score := FindWinner(moves, boards)
    if final_score != 4512 {
        t.Fatalf("Expected Final score: 4512\nGot final score: %d\n\n", final_score)
    }

    moves, boards, err = LoadData("input.txt")
    if err != nil {
        t.Fatalf("Error opening data file: %v", err)
    }
    final_score = FindWinner(moves, boards)
    fmt.Printf("Part 1 solution:\n\nFinal score: %d\n\n", final_score)
}

func Test_Day4_Part2(t *testing.T) {
    moves, boards, err := LoadData("test_data.txt")
    if err != nil {
        t.Fatalf("Error opening test file: %v", err)
    }
    final_score := FindLoser(moves, boards)
    if final_score != 1924 {
        t.Fatalf("Expected Final score: 1924\nGot final score: %d\n\n", final_score)
    }

    moves, boards, err = LoadData("input.txt")
    if err != nil {
        t.Fatalf("Error opening data file: %v", err)
    }
    final_score = FindLoser(moves, boards)
    fmt.Printf("Part 2 solution:\n\nFinal score: %d\n\n", final_score)
}

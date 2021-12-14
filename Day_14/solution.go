package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
    solution, err := part1("test_data.txt", 10)
    if err != nil {
        fmt.Printf("Error opening test file: %v", err)
        os.Exit(1)
    }
    if solution != 1588 {
        fmt.Printf("Test error. Expected 1588, got %d\n", solution)
        os.Exit(1)
    }
    solution, err = part1("input.txt", 10)
    if err != nil {
        fmt.Printf("Error opening test file: %v", err)
        os.Exit(1)
    }
    fmt.Printf("Part 1 solution\n\n%d\n\n", solution)


    solution, err = part1("test_data.txt", 40)
    if err != nil {
        fmt.Printf("Error opening test file: %v", err)
        os.Exit(1)
    }
    if solution != 2188189693529 {
        fmt.Printf("Test error. Expected 2188189693529, got %d\n", solution)
        os.Exit(1)
    }
    solution, err = part1("input.txt", 40)
    if err != nil {
        fmt.Printf("Error opening test file: %v", err)
        os.Exit(1)
    }
    fmt.Printf("Part 2 solution\n\n%d\n\n", solution)
}

func part1(filename string, iterations int) (int, error) {
    polymer, rules, err := LoadData(filename)
    if err != nil {
        return 0, err
    }
    rules_map := make_rules(rules)
    pairs := make_pairs(polymer)
    letters := make_letters(polymer)
    for i := 0; i < iterations; i++ {
        pairs, letters = pair_insertion(pairs, letters, rules_map)
    }
    return find_diff(letters), nil
}

func find_diff(letters map[rune]int) int {
    max_val := 0
    min_val := int(^uint(0) >> 1)
    for _, v := range letters {
        if v < min_val {
            min_val = v
        }
        if v > max_val {
            max_val = v
        }
    }
    return (max_val - min_val)
}

func pair_insertion(polymer_pairs map[string]int, letters map[rune]int, rules map[string][]string) (map[string]int, map[rune] int) {
    out := make(map[string]int)
    for k, w := range polymer_pairs {
        if new, ok := rules[k]; ok {
            for _, v := range new {
                if _, ok := out[v]; ok {
                    out[v] += w
                } else {
                    out[v] = w
                }
            }
            newl := rune(new[0][1])
            if _, ok := letters[newl]; ok {
                letters[newl] += w
            } else {
                letters[newl] = w
            }
        }
    }
    return out, letters
}

func make_letters(polymer string) map[rune]int {
    out := make(map[rune]int)
    for _, v := range polymer {
        if _, ok := out[v]; ok {
            out[v]++
        } else {
            out[v] = 1
        }
    }
    return out
}
func make_pairs(polymer string) map[string]int {
    out := make(map[string]int)
    for i := 0; i < len(polymer)-1; i++ {
        pair := polymer[i:i+2]
        if _, ok := out[pair]; ok {
            out[pair]++
        } else {
            out[pair] = 1
        }
    }
    return out
}

func make_rules(data []string) map[string][]string {
    rules := make(map[string][]string)
    for _, v := range data {
        tmp := strings.Split(v, " -> ")
        rules[tmp[0]] = []string{string(tmp[0][0]) + tmp[1], tmp[1] + string(tmp[0][1])}
    }
    return rules
}

func LoadData(filename string) (string, []string, error) {
    fileBytes, err := ioutil.ReadFile(filename)

    if err != nil {
        return "", nil, err
    }
    raw_data := strings.Split(strings.Trim(string(fileBytes),"\n"), "\n")
    return raw_data[0], raw_data[2:], nil
}

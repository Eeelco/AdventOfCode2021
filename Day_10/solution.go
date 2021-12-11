package day10

import (
	"io/ioutil"
	"sort"
	"strings"
)

func FixIncomplete(data []string) int {
    comp_map := map[rune]rune{'[':']', '(':')', '{':'}', '<':'>'}
    point_map := map[rune]int{'}':3, ']':2, ')':1,'>':4}
    var scores []int
    for _, v := range data {
        corr, rest := FindCorruptChar(v)
        if corr != '!' {
            continue
        } else {
            total_score := 0
            for _, w := range reverse(rest) {
                total_score *= 5
                total_score += point_map[comp_map[w]]
            }
            scores = append(scores, total_score)
        }
    }
    sort.Ints(scores)
    return scores[len(scores)/2]
}

func FindCorruptChar(cmd string) (rune, []rune) {
    var open_braces []rune
    matches := map[rune]rune{'}':'{', ']':'[', ')':'(','>':'<'}
    for _, v := range cmd {
        switch v {
        case '(', '[','{','<':
            open_braces =append(open_braces, v)
        default:
            if open_braces[len(open_braces)-1] == matches[v] {
                open_braces = open_braces[:len(open_braces)-1]
            } else {
                return v, []rune{}
            }
        }
    }
    return '!', open_braces
}

func CorruptedErrors(data []string) int {
    costs := map[rune]int{
        ')':3, ']': 57, '}':1197, '>':25137, '!': 0,
    }
    // tot_cost := 0
    cost_chan := make(chan int, len(data))
    for _, v := range data {
        go func(v string) {
            c, _ := FindCorruptChar(v)
            cost_chan <- costs[c]
        } (v);
        // c, _ := FindCorruptChar(v)
        // tot_cost += costs[c]
    }
    tot_cost := 0
    for i := 0; i < len(data); i++ {tot_cost += <-cost_chan}
    return tot_cost
}

func reverse(numbers []rune) []rune {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func LoadData(filename string) ([]string, error) {
    fileBytes, err := ioutil.ReadFile(filename)

    if err != nil {
        return nil, err
    }
    raw_data := strings.Split(strings.Trim(string(fileBytes),"\n"), "\n")
    return raw_data, nil
}

package day08

import (
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

func IntPow(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
    return len(s)
}

func SortString(s string) string {
    r := []rune(s)
    sort.Sort(sortRunes(r))
    return string(r)
}


func Difference(a string, b string) string {
    output := ""
    for _, v := range a {
        if !strings.Contains(b, string(v)) {
            output += string(v)
        }
    }
    for _, v := range b {
        if !strings.Contains(a, string(v)) {
            output += string(v)
        }
    }
    return output
}

func Union(a string, b string) string {
    var x string
    var y string
    if len(a) > len(b) {
        x = a
        y = b
    } else {
        x = b
        y = a
    }
    for _, v := range x {
        if !strings.Contains(y, string(v)) {
            y += string(v)
        }
    } 
    return y
}

func Intersect(a string, b string) string {
    var x string
    var y string
    if len(a) > len(b) {
        x = a
        y = b
    } else {
        x = b
        y = a
    }
    output := ""
    for _, v := range x {
        if strings.Contains(y, string(v)) {
            output += string(v)
        }
    }
    return output
}

func SplitLines(data []string) ([][]string, [][]string) {
    var digits [][]string
    var output [][] string
    for _, v := range data {
        tmp := strings.Split(v, "|")
        digits = append(digits, strings.Fields(tmp[0]))
        output = append(output, strings.Fields(tmp[1]))
    }
    return digits, output
}

func LoadData(filename string) ([]string, error) {
    fileBytes, err := ioutil.ReadFile(filename)

    if err != nil {
        return nil, err
    }
    raw_data := strings.Split(strings.Trim(string(fileBytes),"\n"), "\n")
    return raw_data, nil
}

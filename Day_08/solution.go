package day08

import (
	"strings"
)

func FindNumber(input []string, digits map[string]int) int {
    nr := 0
    for i, v := range input {
        nr += digits[SortString(v)] * IntPow(10,3-i)
    }
    return nr
}

func DeduceDigits(digits []string) map[string]int {
    output := make(map[int]string)
    for _, v := range digits {
        if len(v) == 2 {
            output[1] = SortString(v)
        } else if len(v) == 3 {
            output[7] = SortString(v)
        } else if len(v) == 4 {
            output[4] = SortString(v)
        } else if len(v) == 7 {
            output[8] = SortString(v)
        }
    }
    for _, v := range digits {
        if len(v) == 6 {
            if len(Union(v,output[1])) == 7 {
                output[6] = SortString(v)
            } else if len(Union(v,output[4])) == 6 {
                output[9] = SortString(v)
            } else {
                output[0] = SortString(v)
            }
        }
    }
    for _,v := range digits {
        if len(v) == 5 && SortString(Union(v, output[6])) == output[6] {
            output[5] = SortString(v)
            break
        }
    }
    for _,v := range digits {
        if len(v) == 5 {
            tmp := Difference(v, output[5])
            if len(tmp) == 2 {
                output[3] = SortString(v)
            } else if len(tmp) == 4 {
                output[2] = SortString(v)
            }
        }
    }
    reversed_map := make(map[string]int)
    for k,v := range output {
        reversed_map[v] = k
    }
    return reversed_map
}

func Count_1478(data []string) int {
    count := 0
    for _, v := range data {
        tmp := strings.Split(v, "|")
        for _,x := range strings.Fields(tmp[1]) {
            if len(x)==2 || len(x)==3 || len(x)==4 || len(x)== 7 {
                count++
            }
        }
    }
    return count
}


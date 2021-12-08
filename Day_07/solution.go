package day07

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func GetBestMedian(data []int) int {
	sort.Ints(data)
	min, max := data[0], data[len(data)-1]
	best := int(^uint(0) >> 1)
	for v := min; v <= max; v++ {
		cost := WeightedDistances(data, v)
		if cost < best {
			best = cost
		} else if cost > best {
			break
		}
	}
	return best
}

func WeightedDistances(data []int, median int) int {
	c := 0
	for _, v := range data {
		n := v - median
		if n < 0 {
			n = -n
		}
		c += n * (n + 1) / 2
	}
	return c
}

func CalcDistances(data []int) int {
	median := GetMedian(data)
	dist_sum := 0
	for _, v := range data {
		tmp := v - median
		if tmp < 0 {
			tmp = -tmp
		}
		dist_sum += tmp
	}
	return dist_sum
}

func GetMedian(data []int) int {
	sort.Ints(data)

	middle := len(data) / 2

	if len(data)%2 == 1 {
		return data[middle]
	} else {
		return (data[middle-1] + data[middle]) / 2
	}
}

func ReadFile(filename string) ([]int, error) {
	fileBytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	stringData := strings.Split(strings.Trim(string(fileBytes), "\n"), ",")

	intData := make([]int, len(stringData))
	for i, v := range stringData {
		intData[i], _ = strconv.Atoi(v)
	}
	return intData, nil
}

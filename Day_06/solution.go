package day06

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func RunSimulations(data []int, days int) int {
	day_list := make([]int, 9)
	for _, v := range data {
		day_list[v] += 1
	}
	for i := 0; i < days; i++ {
		fmt.Printf("After %d days: %v\n", i, day_list)
		spawning := day_list[0]
		day_list = day_list[1:]
		day_list = append(day_list, spawning)
		day_list[6] += spawning
	}
	c := 0
	for _, v := range day_list {
		c += v
	}
	return c
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

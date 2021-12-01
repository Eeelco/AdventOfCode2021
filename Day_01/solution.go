package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := ReadFile("input_1.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	increaseCount := 0

	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			increaseCount++
		}
	}
	fmt.Printf("Number of increases: %d\n", increaseCount)
	//----------------------------------------------------------------------
	totalWindowCount := 0

	lastWindow := 0
	for j := 0; j < 3; j++ {
		lastWindow += data[j]
	}

	for i := 1; i < len(data)-3; i++ {
		nextWindow := 0
		for j := 0; j < 3; j++ {
			nextWindow += data[i+j]
		}
		if nextWindow > lastWindow {
			totalWindowCount++
		}
		lastWindow = nextWindow
	}
	fmt.Printf("Number of windowed increases: %d\n", totalWindowCount)
}

func ReadFile(filename string) ([]int, error) {
	fileBytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	stringData := strings.Split(string(fileBytes), "\n")

	intData := make([]int, len(stringData))
	for i, v := range stringData {
		intData[i], _ = strconv.Atoi(v)
	}
	return intData, nil
}

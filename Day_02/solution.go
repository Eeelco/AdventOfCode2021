package day02

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func FindEndpoint(directions *[][]string, start_x int, start_y int) (int, int) {
	for _, v := range *directions {
		delta, _ := strconv.Atoi(v[1])
		switch v[0] {
		case "down":
			start_y += delta
		case "up":
			start_y -= delta
		case "forward":
			start_x += delta
		case "backward":
			start_x -= delta
		}
	}
	return start_x, start_y
}

func AimEndpoint(directions *[][]string, start_x int, start_y int, start_aim int) (int, int) {
	for _, v := range *directions {
		delta, _ := strconv.Atoi(v[1])
		switch v[0] {
		case "down":
			start_aim += delta
		case "up":
			start_aim -= delta
		case "forward":
			start_x += delta
			start_y += start_aim * delta
		case "backward":
			start_x -= delta
			start_y -= start_aim * delta
		}
	}
	return start_x, start_y
}

func ReadFile(filename string) ([][]string, error) {
	fileBytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	stringData := strings.Split(strings.Trim(string(fileBytes), "\n"), "\n")
	out := make([][]string, len(stringData))
	for i, v := range stringData {
		out[i] = strings.Split(v, " ")
	}
	return out, nil
}

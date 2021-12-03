package day03

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func CalcRates(data []string) (int64, int64) {
	var epsilon_bits string
	var gamma_bits string
	for i := 0; i < len(data[0]); i++ {
		zero_count := 0
		one_count := 0
		for j := 0; j < len(data); j++ {
			if data[j][i] == '0' {
				zero_count++
			} else {
				one_count++
			}
		}
		if one_count > zero_count {
			gamma_bits += "1"
			epsilon_bits += "0"
		} else {
			gamma_bits += "0"
			epsilon_bits += "1"
		}
	}
	epsilon, _ := strconv.ParseInt(epsilon_bits, 2, 64)
	gamma, _ := strconv.ParseInt(gamma_bits, 2, 64)
	return epsilon, gamma
}

func ReadFile(filename string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return strings.Split(strings.Trim(string(fileBytes), "\n"), "\n"), nil
}

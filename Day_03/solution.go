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

func GeneratorRatings(data []string) (int64, int64) {
	var co_bits string
	var oxygen_bits string

	for i := 0; i < len(data[0]); i++ {
		oxygen_one_count := 0
		co_one_count := 0
		oxygen_zero_count := 0
		co_zero_count := 0
		for j := 0; j < len(data); j++ {
			if strings.HasPrefix(data[j], co_bits) {
				if data[j][i] == '1' {
					co_one_count++
				} else {
					co_zero_count++
				}
			}
			if strings.HasPrefix(data[j], oxygen_bits) {
				if data[j][i] == '1' {
					oxygen_one_count++
				} else {
					oxygen_zero_count++
				}
			}
		}
		if oxygen_one_count >= oxygen_zero_count {
			oxygen_bits += "1"
		} else {
			oxygen_bits += "0"
		}
		if (co_one_count < co_zero_count && co_one_count > 0) || co_zero_count == 0 {
			co_bits += "1"
		} else {
			co_bits += "0"
		}
	}

	co, _ := strconv.ParseInt(co_bits, 2, 64)
	oxygen, _ := strconv.ParseInt(oxygen_bits, 2, 64)
	return co, oxygen
}

func ReadFile(filename string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return strings.Split(strings.Trim(string(fileBytes), "\n"), "\n"), nil
}

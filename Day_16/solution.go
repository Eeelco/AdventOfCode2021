package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Packet struct {
	version  uint64
	type_ID  uint64
	value    uint64
	children []Packet
}

func main() {
	TestPartOne()
	data, err := LoadData("input.txt")
	if err != nil {
		fmt.Println("Error loading input")
		os.Exit(2)
	}
	pack, _ := ParsePacket(TranslateData(data))
	sum := SumVersions(pack)
	fmt.Printf("Part 1 solution\n\n%d\n\n", sum)

	TestPartTwo()
	result := EvalExpression(pack)
	fmt.Printf("Part 2 solution\n\n%d\n\n", result)
}


func EvalExpression(pack Packet) uint64 {
	switch pack.type_ID {
	case 4:
		return pack.value
	case 0:
		v := uint64(0)
		for _, p := range pack.children {v += EvalExpression(p)}
		return v
	case 1:
		v := uint64(1)
		for _, p := range pack.children {v *= EvalExpression(p)}
		return v
	case 2:
		min := ^uint64(0)
		for _, p := range pack.children { 
			tmp := EvalExpression(p)
			if tmp < min {min = tmp}
		}
		return min
	case 3:
		max := uint64(0)
		for _, p := range pack.children { 
			tmp := EvalExpression(p)
			if tmp > max {max = tmp}
		}
		return max
	case 5:
		if EvalExpression(pack.children[0]) > EvalExpression(pack.children[1]) {
			return 1
		} else {
			return 0
		}
	case 6:
		if EvalExpression(pack.children[0]) < EvalExpression(pack.children[1]) {
			return 1
		} else {
			return 0
		}
	case 7:
		if EvalExpression(pack.children[0]) == EvalExpression(pack.children[1]) {
			return 1
		} else {
			return 0
		}
	default:
		fmt.Printf("Error parsing packet")
		return 0
	}
}

func ParsePacket(data string) (Packet, string) {
	version, _ := strconv.ParseUint(data[0:3], 2, 3)
	type_ID, _ := strconv.ParseUint(data[3:6], 2, 3)
	pack := Packet{version, type_ID, 0, nil}
	data = data[6:]
	if type_ID == 4 {
		return ParseValue(pack, data)
	} else {
		length_type := data[0]
		data = data[1:]
		if length_type == '0' {
			packet_length, _ := strconv.ParseUint(data[:15], 2, 15)
			data = data[15:]
			sub_packets := data[:packet_length]
			data = data[packet_length:]
			for len(sub_packets) > 0 {
				newpack, rest := ParsePacket(sub_packets)
				sub_packets = rest
				pack.children = append(pack.children, newpack)
			}
			return pack, data
		} else {
			child_count, _ := strconv.ParseUint(data[:11], 2, 11)
			data = data[11:]
			for uint64(len(pack.children)) < child_count {
				new_child, rest := ParsePacket(data)
				data = rest
				pack.children = append(pack.children, new_child)
			}
			return pack, data
		}
	}
}

func ParseValue(pack Packet, data string) (Packet, string) {
	val := ""
	for {
		next := data[:5]
		data = data[5:]
		val += next[1:]
		if next[0] == '0' {
			pack.value, _ = strconv.ParseUint(val, 2, len(val))
			return pack, data
		}
	}
}

func TranslateData(hex string) (out string) {
	for _, v := range hex {
		a, _ := strconv.ParseUint(string(v), 16, 4)
		b := strconv.FormatInt(int64(a), 2)
		out += fmt.Sprintf("%04v", b)
	}
	return
}

func SumVersions(pack Packet) (sum uint64) {
	sum += pack.version
	for _, c := range pack.children {
		sum += SumVersions(c)
	}
	return sum
}

func LoadData(filename string) (string, error) {
	fileBytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return "", err
	}
	return strings.Trim(string(fileBytes),"\n") , nil
}

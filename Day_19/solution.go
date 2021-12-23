package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	test_scanners, err := LoadData("test_data.txt")
	if err != nil {
		fmt.Print("Could not open test file")
		os.Exit(1)
	}
	scanner_positions, points := unite_all(test_scanners)
	fmt.Printf("Nr. of scanners: %d\n\nNr. of beacons: %d\n\nMax. scanner distance: %v\n\n", len(scanner_positions), len(points), max_scanner_distance(scanner_positions))
	real_scanners, err := LoadData("input.txt")
	if err != nil {
		fmt.Print("Could not open test file")
		os.Exit(1)
	}
	scanner_positions, points = unite_all(real_scanners)
	fmt.Printf("Solution part 1:\n\n%d\n\nSolution part 2:\n\n%v\n\n", len(points), max_scanner_distance(scanner_positions))
}

func max_scanner_distance(scanners []Point) float64 {
	max_dist := 0.0
	for i := 0; i < len(scanners); i++ {
		for j := 0; j < i; j++ {
			dist := manhattan_dist(scanners[i], scanners[j])
			if dist > max_dist {
				max_dist = dist
			}
		}
	}
	return max_dist
}

func unite_all(scanners [][]Point) ([]Point, []Point) {
	fixed_positions := scanners[0]
	scanners = scanners[1:]
	scanner_positions := []Point{{0, 0, 0}}
	for len(scanners) > 0 {
		new_scanner_pos, new_pos, s := find_union(fixed_positions, scanners)
		fixed_positions = join(fixed_positions, new_pos)
		scanner_positions = append(scanner_positions, new_scanner_pos)
		scanners = s
	}
	return scanner_positions, fixed_positions
}

func find_union(reference []Point, scanners [][]Point) (Point, []Point, [][]Point) {
	for j := 0; j < len(scanners); j++ {
		check, p1, p2 := do_overlap(reference, scanners[j])
		if check {
			var line Line
			var dist float64
			for line, dist = range p1 {
				break
			}
			vec := sub(line.first, line.second)
			corresponding_pair := getKey(p2, dist)
			rots1, rots2 := get_all_rots(corresponding_pair.first), get_all_rots(corresponding_pair.second)
			var i int
			var vec_forward, vec_backward Point
			for i = 0; i < len(rots1); i++ {
				vec_forward, vec_backward = sub(rots1[i], rots2[i]), sub(rots2[i], rots1[i])
				if vec == vec_forward || vec == vec_backward {
					break
				}
			}
			var anchor_point Point
			if vec == vec_forward {
				anchor_point = rots1[i]
			} else {
				anchor_point = rots2[i]
			}
			second_beacon_pos := sub(line.first, anchor_point)
			var corrected_positions []Point
			for _, p := range rotate_all_points(scanners[j])[i] {
				corrected_positions = append(corrected_positions, add(p, second_beacon_pos))
			}
			return second_beacon_pos, corrected_positions, append(scanners[:j], scanners[j+1:]...)
		}
	}
	fmt.Printf(("No overlap found"))
	os.Exit(2)
	return Point{}, nil, nil
}

func do_overlap(sc1 []Point, sc2 []Point) (bool, map[Line]float64, map[Line]float64) {
	s1_dists := get_all_dists(sc1)
	s2_dists := get_all_dists(sc2)
	intersect, l1, l2 := Intersect(s1_dists, s2_dists)
	if len(intersect) >= 66 {
		return true, l1, l2
	}
	return false, nil, nil
}

func get_all_dists(points []Point) map[Line]float64 {
	dists := make(map[Line]float64)
	for i := 0; i < len(points); i++ {
		for j := 0; j < i; j++ {
			dists[Line{points[i], points[j]}] = dist(points[i], points[j])
		}
	}
	return dists
}

func rotate_all_points(points []Point) [][]Point {
	out := make([][]Point, 24)
	for _, p := range points {
		for j, r := range get_all_rots(p) {
			out[j] = append(out[j], r)
		}
	}
	return out
}

func get_all_rots(point Point) (out []Point) {
	p := point
	for i := 0; i < 2; i++ {
		out = append(out, p)
		p = z_90deg_rot(p)
		out = append(out, p)
		p = z_90deg_rot(p)
		out = append(out, p)
		p = z_90deg_rot(p)
		out = append(out, p)
		p = z_90deg_rot(p)
		p = x_90deg_rot(x_90deg_rot(p))
	}

	p = x_90deg_rot(p)
	for j := 0; j < 2; j++ {
		for i := 0; i < 4; i++ {
			out = append(out, p)
			p = z_90deg_rot(p)
		}
		p = x_90deg_rot(x_90deg_rot(p))
	}
	p = x_90deg_rot(x_90deg_rot(x_90deg_rot(p)))

	p = y_90deg_rot(p)
	for i := 0; i < 4; i++ {
		out = append(out, p)
		p = z_90deg_rot(p)
	}
	p = y_90deg_rot(y_90deg_rot(p))
	for i := 0; i < 4; i++ {
		out = append(out, p)
		p = z_90deg_rot(p)
	}
	return
}

func join(a []Point, b []Point) (c []Point) {
	var x []Point
	var y []Point
	if len(a) > len(b) {
		x = a
		y = b
	} else {
		x = b
		y = a
	}
	for _, v := range x {
		if !Pcontains(y, v) {
			y = append(y, v)
		}
	}
	return y

}
func Pcontains(a []Point, p Point) bool {
	for _, v := range a {
		if v == p {
			return true
		}
	}
	return false
}

func LoadData(filename string) ([][]Point, error) {
	fileBytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.Trim(string(fileBytes), "\n"), "\n")
	var outdata [][]Point
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		if l[1] == '-' {
			outdata = append(outdata, []Point{})
		} else {
			ps := strings.Split(l, ",")
			x, _ := strconv.ParseFloat(ps[0], 64)
			y, _ := strconv.ParseFloat(ps[1], 64)
			z, _ := strconv.ParseFloat(ps[2], 64)
			outdata[len(outdata)-1] = append(outdata[len(outdata)-1], Point{x, y, z})
		}
	}
	return outdata, nil
}

package main

import "math"

type Point struct {
	x float64
	y float64
	z float64
}
type Line struct {
	first  Point
	second Point
}

func x_90deg_rot(point Point) Point {
	return Point{point.x, -point.z, point.y}
}
func y_90deg_rot(point Point) Point {
	return Point{point.z, point.y, -point.x}
}
func z_90deg_rot(point Point) Point {
	return Point{-point.y, point.x, point.z}
}

func sub(a Point, b Point) Point {
	return Point{a.x - b.x, a.y - b.y, a.z - b.z}
}
func add(a Point, b Point) Point {
	return Point{a.x + b.x, a.y + b.y, a.z + b.z}
}

func dist(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2) + math.Pow(p1.z-p2.z, 2))
}
func manhattan_dist(p1 Point, p2 Point) float64 {
	return math.Abs(p1.x-p2.x) + math.Abs(p1.y-p2.y) + math.Abs(p1.z-p2.z)
}

func Intersect(a map[Line]float64, b map[Line]float64) ([]float64, map[Line]float64, map[Line]float64) {
	set1lines := make(map[Line]float64)
	set2lines := make(map[Line]float64)
	var out []float64
	for k, v := range a {
		if contains(b, v) {
			out = append(out, v)
			set1lines[k] = v
			set2lines[getKey(b, v)] = v
		}
	}
	return out, set1lines, set2lines
}
func contains(s map[Line]float64, str float64) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func getKey(s map[Line]float64, t float64) Line {
	for k, v := range s {
		if v == t {
			return k
		}
	}
	return Line{}
}

package main

import "math"

type Point []float64

func x_90deg_rot(point Point) Point {
    return Point{point[0], -point[2], point[1]}
}
func y_90deg_rot(point Point) Point {
    return Point{point[2], point[1], -point[0]}
}
func z_90deg_rot(point Point) Point {
    return Point{-point[1], point[0], point[2]}
}

func sub(a Point, b Point) Point {
    return Point{a[0]-b[0],a[1]-b[1],a[2]-b[2]}
}
func add(a Point, b Point) Point {
    return Point{a[0]+b[0],a[1]+b[1],a[2]+b[2]}
}

func dist(p1 Point, p2 Point) float64 {
    return math.Sqrt(math.Pow(p1[0] - p2[0],2) + math.Pow(p1[1] - p2[1],2) + math.Pow(p1[2] - p2[2],2))
}

func Union(a []float64, b []float64) []float64 {
    var x []float64
    var y []float64
    if len(a) > len(b) {
        x = a
        y = b
    } else {
        x = b
        y = a
    }
    for _, v := range x {
        if !contains(y, v) {
            y = append(y, v)
        }
    } 
    return y
}
func contains(s []float64, str float64) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

package main

import (
	"fmt"
	"os"
)

func unite_all(scanners [][]Point) {
    var scanner_positions []Point
    for len(scanners) > 1 {
        scanner_pos, corrected_pos, rest := find_union(scanners)
        scanner_positions = append(scanner_positions, scanner_pos)
        scanners = rest
        scanners[0] = join(scanners[0], corrected_pos)
    }
}

func find_union(scanners [][]Point) (Point, []Point, [][]Point){
    s1 := scanners[0]
    for j := 1; j < len(scanners); j++ {
        check, p1, p2 := do_overlap(s1, scanners[j])
        if check {
            var line Line;var dist float64; for line, dist = range p1 { break }
            vec := sub(line.first , line.second)
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
            if vec == vec_forward {anchor_point = rots1[i]} else {anchor_point = rots2[i]}
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
            dists[Line{points[i], points[j]}] = dist(points[i],points[j])
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
    for i := 0; i < 4; i++ {
        out = append(out, p)
        p = x_90deg_rot(p)
        out = append(out, p)
        p = x_90deg_rot(x_90deg_rot(p))
        out = append(out, p)
        p = x_90deg_rot(x_90deg_rot(x_90deg_rot(p)))
        out = append(out, p)
        p = z_90deg_rot(p)
    }
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

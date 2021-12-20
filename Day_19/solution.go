package main

func do_overlap(sc1 []Point, sc2 []Point) bool {
    s1_dists := get_all_dists(sc1)
    s2_dists := get_all_dists(sc2)
    if len(Union(s1_dists, s2_dists)) >= 66 {
        return true
    }
    return false
}

func get_all_dists(points []Point) (dists []float64) {
    for i := 0; i < len(points); i++ {
        for j := 0; j < i; j++ {
            dists = append(dists, dist(points[i], points[j]))
        }
    }
    return dists
}

func check_overlap(base_scanner []Point, possible_oris [][]Point) []Point {

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

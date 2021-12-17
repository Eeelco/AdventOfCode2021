package main

import (
	"fmt"
	"math"
)

type Area struct {
    xmin int
    xmax int
    ymin int
    ymax int
}

func (a *Area) contains(x int, y int) bool {
    return a.xmin <= x && x <= a.xmax && a.ymin <= y && y <= a.ymax
}
func (a *Area) misses(x int, y int) bool {
    return x > a.xmax || y < a.ymin
}

func main() {
    test_target := Area{20, 30, -10, -5}
    ymax := find_best_parabola(test_target)
    if ymax != 45 {
        fmt.Printf("Test error. Expected 45, got %d\n", ymax)
    }
    real_target := Area{230, 283, -107, -57}
    ymax = find_best_parabola(real_target)
    fmt.Printf("Part 1 solution: %d\n\n", ymax)

    all_hits := find_all_parabolas(test_target)
    if all_hits != 112 {
        fmt.Printf("Test error. Expected 112, got %d\n", all_hits)
    }
    all_hits = find_all_parabolas(real_target)
    fmt.Printf("Part 2 solution: %d\n\n", all_hits)
}

func find_best_parabola(target Area) int {
    min_vx := int(0.5 * (math.Sqrt(8 * float64(target.xmin)) - 1))
    max_vx := target.xmax
    best_ymax := 0
    for vx:= min_vx; vx <= max_vx; vx++ {
        for vy := target.ymin; vy <= max_vx; vy++ {
            ymax := parabola(vx, vy, target)
            if ymax > best_ymax {
                best_ymax = ymax
            }
        }
    }
    return best_ymax
}

func find_all_parabolas(target Area) int {
    min_vx := int(0.5 * (math.Sqrt(8 * float64(target.xmin)) - 1))
    max_vx := target.xmax
    parabola_count := 0
    for vx:= min_vx; vx <= max_vx; vx++ {
        for vy := target.ymin; vy <= max_vx; vy++ {
            ymax := parabola(vx, vy, target)
            if ymax > -1 {
                parabola_count++
            }
        }
    }
    return parabola_count
}

func parabola(vx int, vy int, target Area) int {
    x, y := 0, 0
    y_max := 0
    for {
        x += vx
        y += vy
        if y > y_max {
            y_max = y
        }
        if target.contains(x, y) {
            return y_max
        } else if target.misses(x, y) {
            return -1
        }
        vx = int(math.Max(0, float64(vx-1)))
        vy -= 1
    }
}

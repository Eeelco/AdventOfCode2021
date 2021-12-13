package day13

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
    x int
    y int
}
func GetSolution(points []Point, folds []string) {
    for _, v := range folds {
        points = PerformFold(points, v)
    }
    render(points)
}

func render(points []Point) {
    max_x, max_y := 0, 0
    for _, p := range points {
        if p.x > max_x {max_x = p.x}
        if p.y > max_y {max_y = p.y}
    }
    output := make([][]byte, max_y+1)
    for i:=0; i <=max_y; i++ {
        output[i] = make([]byte, max_x + 1)
        for j:= 0; j <= max_x; j++ {output[i][j] = '.'}
    }
    for _, p := range points {
        output[p.y][p.x] = '#'
    }
    for _, l := range output {
        fmt.Println(string(l))
    }
}

func PerformFold(points []Point, folds string) []Point {
    all_points := make(map[Point]bool)
    re := regexp.MustCompile("([x|y])=(\\d+)")
    data := re.FindStringSubmatch(folds)
    line, _ := strconv.Atoi(data[2])
    for _, p := range points {
        if data[1] == "x" {
            if p.x > line {
                p.x = line - (p.x - line)
            }
        } else {
            if p.y > line {
                p.y = line - (p.y - line)
            }
        }
        if _, ok := all_points[p]; !ok {
            all_points[p] = true
        }
    }
    output := make([]Point, len(all_points))
    i := 0
    for k := range all_points {
        output[i] = k
        i++
    }
    return output
}

func LoadData(filename string) ([]Point, []string, error) {
    fileBytes, err := ioutil.ReadFile(filename)

    if err != nil {
        return nil, nil, err
    }
    var initial_points []Point
    var folds []string
    raw_data := strings.Split(strings.Trim(string(fileBytes),"\n"), "\n")
    for _, v := range raw_data {
        if !strings.Contains(v, "fold") && len(v) > 0 {
            tmp := strings.Split(v, ",")
            a,_ := strconv.Atoi(tmp[0])
            b,_ := strconv.Atoi(tmp[1])
            initial_points = append(initial_points, Point{a,b})
        } else if len(v) > 0 {
            folds = append(folds, v)
        }
    }
    return initial_points, folds, nil
}

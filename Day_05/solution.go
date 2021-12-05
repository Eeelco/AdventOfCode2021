package day05

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func GetResult(lines [][][]int, no_diagonals bool) int {
    line_points := make([][][]int, len(lines))
    for i,v := range lines {
        line_points[i] = Calcline(v, no_diagonals)
    }
    return CountIntersections(line_points)
}

func CountIntersections(lines [][][]int) int {
    data := make(map[string]int)
    for _, v := range lines {
        for _,w := range v {
            key := strconv.Itoa(w[0]) + "," + strconv.Itoa(w[1])
            _,ok := data[key]
            if ok {
                data[key]++
            } else {
                data[key] = 1
            }
        }
    }
    count := 0
    for _, v := range(data) {
        if v > 1 {
            count++
        }
    }
    return count
}

func Calcline(points [][]int, no_diagonals bool) [][]int {
    p1 := points[0]
    p2 := points[1]
    if no_diagonals && (p1[1] != p2[1] && p1[0] != p2[0]) {
        return [][]int{}
    }
    dists := []float64{math.Abs(float64(p1[0]-p2[0])),math.Abs(float64(p1[1]-p2[1]))}
    dist := int(math.Max(dists[0],dists[1]))
    output := make([][]int, dist+1)
    for i:= 0; i < dist+1; i++ {
        output[i]= []int{p1[0]+i*(p2[0]-p1[0])/dist, p1[1] + i*(p2[1]-p1[1])/dist}
    }
    return output
}

func LoadData(filename string) ([][][]int, error) {
    fileBytes, err := ioutil.ReadFile(filename)

    if err != nil {
        return nil, err
    }
    raw_data := strings.Split(strings.Trim(string(fileBytes),"\n"), "\n")
    output := make([][][]int, len(raw_data))
    for i, v := range raw_data {
        tmp := strings.Split(v, " -> ")
        p1 := strings.Split(tmp[0], ",")
        p2 := strings.Split(tmp[1], ",")
        x1, _ := strconv.Atoi(p1[0])
        y1, _ := strconv.Atoi(p1[1])
        x2, _ := strconv.Atoi(p2[0])
        y2, _ := strconv.Atoi(p2[1])

        output[i] = [][]int{{x1, y1},{x2, y2}}
    }
    return output, nil
}

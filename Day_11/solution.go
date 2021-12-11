package day11

import (
	"io/ioutil"
	"strings"
)

func Synchronization(filename string) (int, error) {
    data, err := LoadData(filename)
    if err != nil {
        return 0, err
    }
    n_octopi := len(data) * len(data)
    for i := 1; ; i++{
        flashes := TimeStep(&data)
        if flashes == n_octopi {
            return i, nil
        }
    }
}

func Simulate(filename string, n_steps int) (int, error) {
    data, err := LoadData(filename)
    if err != nil {
        return 0, err
    }
    flashes := 0
    for i := 0; i < n_steps; i++ {
        flashes += TimeStep(&data)
    }
    return flashes, nil
}

func TimeStep(grid *[][]int) int {
    flashes := 0
    var flash_coords [][]int
    var has_flashed [][]int
    for i, v := range *grid {
        for j, _ := range v {
            (*grid)[i][j]++
            if (*grid)[i][j] == 10 {
                flashes++
                flash_coords = append(flash_coords, []int{i,j})
                has_flashed = append(has_flashed, []int{i,j})
            }
        }
    }
    for len(flash_coords) > 0 {
        co := flash_coords[0]
        flash_coords = flash_coords[1:]
        nbs := GetNeighbors(co[0], co[1], len(*grid))
        for _,n := range nbs {
            (*grid)[n[0]][n[1]]++
            if (*grid)[n[0]][n[1]] == 10 {
                flashes++
                // (*grid)[n[0]][n[1]] = 0
                flash_coords = append(flash_coords, n)
                has_flashed = append(has_flashed, n)
            }
        }
    }
    for _,s := range has_flashed {
        (*grid)[s[0]][s[1]] = 0
    }
    return flashes
}

func GetNeighbors(x int, y int, L int) [][]int {
    var output [][]int
    for i := -1; i < 2; i++ {
        for j := -1; j < 2; j++ {
            if i == 0 && j == 0 {continue}
            newx, newy := x+i, y+j
            if 0 <= newx && newx < L && 0 <=newy && newy < L {
                output = append(output, []int{newx, newy})
            }
        }
    }
    return output
}

func LoadData(filename string) ([][]int, error) {
    fileBytes, err := ioutil.ReadFile(filename)

    if err != nil {
        return nil, err
    }
    raw_data := strings.Split(strings.Trim(string(fileBytes),"\n"), "\n")
    int_data := make([][]int, len(raw_data))
    for i, v := range raw_data {
        int_data[i] = make([]int, len(v))
        for j, w := range v {
            int_data[i][j] = int(w-'0')
        }
    }
    return int_data, nil
}

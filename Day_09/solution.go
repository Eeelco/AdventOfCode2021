package day09

import (
	"io/ioutil"
	"sort"
	"strings"
)

func isSafe(data *[][]int, visited *[][]int, i int, j int) bool {
    return i >= 0 && i < len(*data) &&
    j >= 0 && j < len((*data)[0]) &&
    (*data)[i][j] < 9 && (*visited)[i][j] == 0
}

func BFS(data *[][]int, visited *[][]int, si int, sj int) int {
    island_size := 1
    row := []int{-1,0,1,0}
    col := []int{0,-1,0,1}
    queue := make([][]int, 0)
    queue = append(queue, []int{si,sj})
    (*visited)[si][sj] = 1
    for len(queue) > 0 {
        tmp := queue[0]
        queue = queue[1:]
        i,j := tmp[0], tmp[1]
        for k:= 0; k < 4; k++ {
            x,y := i+row[k], j+col[k]
            if isSafe(data, visited, x, y) {
                (*visited)[x][y] = 1
                queue = append(queue, []int{x, y})
                island_size++
            }
        }
    }
    return island_size
}

func find_basins(data [][]int) int {
    var basin_sizes []int
    visited := make([][]int, len(data))
    for i:=0;i<len(data);i++ {
        visited[i] = make([]int, len(data[0]))
    }
    for i, v := range data {
        for j, _ := range v {
            if data[i][j] < 9 && visited[i][j] == 0 {
                basin_sizes = append(basin_sizes, BFS(&data, &visited,i,j))
            }
        }
    }
    sort.Sort(sort.Reverse(sort.IntSlice(basin_sizes)))
    return basin_sizes[0] * basin_sizes[1] * basin_sizes[2]
}

func count_local_minima(data [][]int) int {
    mins := 0
    for i, v := range data {
        for j, w := range v {
            if (i == 0 || data[i][j] < data[i-1][j]) &&
            (j == 0 || data[i][j] < data[i][j-1]) &&
            (i == len(data)-1 || data[i][j] < data[i+1][j]) &&
            (j == len(v) - 1 || data[i][j] < data[i][j+1]) {
                mins += w+1
            }

        }
    }
    return mins
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

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"

	pq "github.com/kyroy/priority-queue"
)

type Node struct {
    X int
    Y int
    Risk int
}

func main() {
    test_data, err := LoadData("test_data.txt")
    if err != nil {
        fmt.Printf("Error loading test data")
        os.Exit(1)
    }
    _, fScore, err := A_star(test_data[0][0], test_data[len(test_data)-1][len(test_data[0])-1], test_data)
    if err != nil {
        fmt.Printf("Could not find path through test data")
        os.Exit(2)
    }
    if int(fScore) != 40 {
        fmt.Printf("Error calculating path. Expected 40, got %v\n", fScore)
        os.Exit(3)
    }
    input_data, err := LoadData("input.txt")
    if err != nil {
        fmt.Printf("Error loading test data")
        os.Exit(1)
    }
    _, fScore, err = A_star(input_data[0][0], input_data[len(input_data)-1][len(input_data[0])-1], input_data)
    if err != nil {
        fmt.Printf("Could not find path through data")
        os.Exit(2)
    }
    fmt.Printf("Part 1 solution: %d\n\n", int(fScore))

    larger_test_data := make_larger_grid(test_data, 5)
    _, fScore, err = A_star(larger_test_data[0][0], larger_test_data[len(larger_test_data)-1][len(larger_test_data[0])-1], larger_test_data)
    if int(fScore) != 315 {
        fmt.Printf("Error calculating path. Expected 40, got %v\n", fScore)
        os.Exit(3)
    }
    larger_input_data := make_larger_grid(input_data, 5)
    _, fScore, err = A_star(larger_input_data[0][0], larger_input_data[len(larger_input_data)-1][len(larger_input_data[0])-1], larger_input_data)
    fmt.Printf("Part 2 solution: %d\n\n", int(fScore))

}

func make_larger_grid(grid [][]Node, times int) [][]Node {
    output := make([][]Node, times * len(grid))
    for i := 0; i < times*len(grid); i++ {
        output[i] = make([]Node, times * len(grid[0]))
        for j := 0; j < times*len(grid[0]); j++ {
            new_val := grid[i%len(grid)][j%len(grid[0])].Risk
            for x := 0; x < i/len(grid) + j/len(grid[0]); x++ {
                new_val++
                if new_val == 10 {new_val = 1}
            }
            output[i][j] = Node{i, j, new_val}
        }
    }
    return output
}

func h(start Node, end Node) float64 {
    return math.Sqrt(math.Pow(float64(start.X)-float64(end.X), 2) + math.Pow(float64(start.Y)-float64(end.Y), 2))
}

func A_star(start Node, end Node, all_nodes[][]Node) ([]Node, float64, error){
    pq := pq.NewPriorityQueue()
    pq.Insert(start, h(start, end))
    came_from := make(map[Node]Node)

    gScore := make(map[Node]int)
    fScore := make(map[Node]float64)
    for i := 0; i < len(all_nodes); i++ {
        for j := 0; j < len(all_nodes[0]); j++ {
            gScore[all_nodes[i][j]] = int(^uint(0)>>1)
            fScore[all_nodes[i][j]] = math.Inf(1)
        }
    }
    gScore[start] = 0
    fScore[start] = h(start, end)

    for pq.Len() > 0 {
        current := pq.PopLowest().(Node)
        if current.X == end.X && current.Y == end.Y {
            return reconstruct_path(came_from, current), fScore[current], nil
        }
        for _, n := range get_neighbors(current, all_nodes) {
            tentative_gScore := gScore[current] + n.Risk
            if tentative_gScore < gScore[n] {
                came_from[n] = current
                gScore[n] = tentative_gScore
                fScore[n] = float64(tentative_gScore) + h(n, end)
                if !Contains(pq, n) {
                    pq.Insert(n, fScore[n])
                }
            }
        }
    }
    return nil, 0, errors.New("No path found")
}

func reconstruct_path(came_from map[Node]Node, current Node) []Node {
    total_path := []Node{current}
    for val, ok := came_from[current]; ok; {
        current = val
        total_path = append([]Node{current}, total_path...)
        val, ok = came_from[current]
    }
    return total_path
}

func Contains(q *pq.PriorityQueue, n Node) bool {
    for i := 0; i < q.Len(); i++ {
        d, _ := q.Get(i)
        if d == n {
            return true
        }
    }
    return false
}

func get_neighbors(n Node, grid [][]Node) (out []Node){
    x, y := n.X, n.Y
    lx, ly := len(grid), len(grid[0])
    if x > 0 { out = append(out, grid[x-1][y])}
    if y > 0 { out = append(out, grid[x][y-1])}
    if x < lx-1 { out = append(out, grid[x+1][y])}
    if y < ly-1 { out = append(out, grid[x][y+1])}

    return
}

func LoadData(filename string) ([][]Node, error) {
    fileBytes, err := ioutil.ReadFile(filename)

    if err != nil {
        return nil, err
    }
    raw_data := strings.Split(strings.Trim(string(fileBytes),"\n"), "\n")
    int_data := make([][]Node, len(raw_data))
    for i, v := range raw_data {
        int_data[i] = make([]Node, len(v))
        for j, w := range v {
            int_data[i][j] = Node{i, j, int((w-'0')) }
        }
    }
    return int_data, nil
}

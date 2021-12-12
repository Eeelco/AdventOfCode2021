package day12

import (
	"io/ioutil"
	"strings"
)

func GetAllPaths(filename string) (int, int, error) {
    connections, err := LoadData(filename)
    if err != nil {
        return 0, 0, err
    }
    nodes := GetNodes(connections)
    graph := NewGraph("start", "end", nodes, connections, false)
    graph_part_2 := NewGraph("start", "end", nodes, connections, true)
    return len(graph.paths), len(graph_part_2.paths), nil
}

func GetNodes(data []string) (node_list []string) {
    for _, v := range data {
        tmp := strings.Split(v, "-")
        a, b := tmp[0], tmp[1]
        if !contains(node_list, a) {
            node_list = append(node_list, a)
        }
        if !contains(node_list, b) {
            node_list = append(node_list, b)
        }
    }
    return
}

func LoadData(filename string) ([]string, error) {
    fileBytes, err := ioutil.ReadFile(filename)

    if err != nil {
        return nil, err
    }
    raw_data := strings.Split(strings.Trim(string(fileBytes),"\n"), "\n")
    return raw_data, nil
}

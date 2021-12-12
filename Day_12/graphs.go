package day12

import "strings"

type Node struct {
    name string
    neighbors []Node
}

type Graph struct {
    start string
    end string
    indices map[string]int
    nodes []Node
    paths [][]string
}

func NewGraph(start string, end string, nodes []string, connections []string, double_visit bool) Graph {
    g := Graph{start, end, map[string]int{}, []Node{}, [][]string{}}
    for i, v := range nodes {
        g.indices[v] = i
        g.nodes = append(g.nodes, Node{v, []Node{}})
    }
    for _, v := range connections {
        tmp := strings.Split(v, "-")
        a, b := tmp[0], tmp[1]
        g.get(a).neighbors = append(g.get(a).neighbors, g.nodes[g.indices[b]])
        g.get(b).neighbors = append(g.get(b).neighbors, g.nodes[g.indices[a]])
    }
    if !double_visit {
        g.DFS_construct_paths()
    } else {
        g.DFS_double_visit()
    }
    return g
}

func (g *Graph) get(node string) *Node {
    return &g.nodes[g.indices[node]]
}

func (g *Graph) DFS_double_visit() {
    g.DFS_double_helper([]string{g.start}, false)
}

func (g *Graph) DFS_double_helper(current_path []string, double_visit bool) {
    current_node := g.get(current_path[len(current_path)-1]) 
    if current_node.name == g.end {
        g.paths = append(g.paths, current_path)
    } else {
        for _, v := range current_node.neighbors {
            if v.name != g.start {
                if (v.name == strings.ToLower(v.name) && contains(current_path, v.name)) {
                    if !double_visit {
                        g.DFS_double_helper(append(current_path, v.name), true)
                    }
                } else {
                    g.DFS_double_helper(append(current_path, v.name), double_visit)
                }
            }
        }
    }
}

func (g *Graph) DFS_construct_paths() {
    g.DFS_helper([]string{g.start})
}

func (g *Graph) DFS_helper(current_path []string) {
    current_node := g.get(current_path[len(current_path)-1]) 
    if current_node.name == g.end {
        g.paths = append(g.paths, current_path)
    } else {
        for _, v := range current_node.neighbors {
            if v.name != g.start && !(v.name == strings.ToLower(v.name) && contains(current_path, v.name)) {
                g.DFS_helper(append(current_path, v.name))
            }
        }
    }
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

type Node struct {
    value int
    depth int
    left *Node
    right *Node
}

func main() {
    a := &Node{-1,0, &Node{9,1,nil,nil}, &Node{1, 1, nil, nil}}
    b := &Node{-1,0, &Node{1,1,nil,nil}, &Node{9, 1, nil, nil}}
    c := add(a, b)
    if magnitude(c) != 129 {
        fmt.Printf("Expected 129, got %d\n", magnitude(c))
        os.Exit(1)
    }
    d := &Node{-1, 0, &Node{9,1,nil,nil}, &Node{8, 1, nil, nil}}
    d = add(d, &Node{1, 0, nil, nil})
    d = add(d, &Node{2, 0, nil, nil})
    d = add(d, &Node{3, 0, nil, nil})
    d = add(d, &Node{4, 0, nil, nil})
    fmt.Println(print_tree(d))
    traverse(d,nil,nil)
    fmt.Println(print_tree(d))
}

func print_tree(n *Node) string {
    if n.left == nil && n.right == nil {
        return strconv.Itoa(n.value)
    }
    l := ""
    r := ""
    if n.left != nil {
        l = print_tree(n.left)
    }
    if n.right != nil {
        r = print_tree(n.right)
    }
    return "[" + l + "," + r + "]"
}

func reduce(n *Node) {
    v := traverse(n,nil,nil)
    for v {
        v = traverse(n,nil,nil)
    }
}

func magnitude(node *Node) int {
    if node.value != -1 {
        return node.value
    } 
    return 3 * magnitude(node.left) + 2 * magnitude(node.right)
}

func inc(nr *Node) {
    nr.depth++
    if nr.left != nil {
        inc(nr.left)
    }
    if nr.right != nil {
        inc(nr.right)
    }
}

func add(number1 *Node, number2 *Node) *Node {
    inc(number1)
    inc(number2)
    return &Node{-1, 0, number1, number2}
}

func get_right_nb(n *Node ) *Node {
    if n.value != -1 {
        return n
    }
    if n.left != nil {
        get_right_nb(n.left)
    }
    if n.right != nil {
        get_right_nb(n.right)
    }
    return nil
}

func get_left_nb(n *Node ) *Node {
    if n.value != -1 {
        return n
    }
    if n.right != nil {
        get_right_nb(n.right)
    }
    if n.left != nil {
        get_right_nb(n.left)
    }
    return nil
}

func traverse(n *Node, left_nb *Node, right_nb *Node) bool {
    if n.left != nil {
        traverse(n.left, left_nb, n.right)
    }

    if n.depth > 3 && n.value == -1 {
        if left_nb != nil {
            get_left_nb(left_nb).value += n.left.value
        }
        if right_nb != nil {
            get_right_nb(right_nb).value += n.right.value
        }
        *n = Node{0, n.depth, nil, nil}
        return true
    }
    if n.value >= 10 {
        lv := n.value / 2
        rv := n.value / 2
        if n.value % 2 == 1{
            rv++
        }
        l := &Node{lv, n.depth + 1, nil, nil}
        r := &Node{rv, n.depth + 1, nil, nil}
        *n = Node{-1, n.depth, l, r}
        return true
    }

    if n.right != nil {
        traverse(n.right, n.left, right_nb)
    }
    return false
}

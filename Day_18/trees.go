package main

import (
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

type Node struct {
	value int
	depth int
	left  *Node
	right *Node
}

func AddNumbers(nrs []string) *Node {
	a := MakeTree(nrs[0])
	// fmt.Println(print_tree(a))
	b := MakeTree(nrs[1])
	// fmt.Println(print_tree(b))
	c := add(a, b)
	reduce(c)
	for i := 2; i < len(nrs); i++ {
		// fmt.Println(print_tree(c))
		c = add(c, MakeTree(nrs[i]))
		reduce(c)
	}
	return c
}

func MakeTree(s string) *Node {
	return TreeFromString(s[1:len(s)-1], 0)
}

func TreeFromString(s string, depth int) *Node {
	root := &Node{-1, depth, nil, nil}
	i := 0
	if unicode.IsDigit(rune(s[0])) {
		tmp := string(s[0])
		i++
		for unicode.IsDigit(rune(s[i])) {
			tmp += string(s[i])
			i++
		}
		a, _ := strconv.Atoi(tmp)
		root.left = &Node{a, depth + 1, nil, nil}
	} else {
		i = findMatchingIndex(s, 0, len(s)-1)
		root.left = TreeFromString(s[1:i], depth+1)
		i++
	}
	i++
	if unicode.IsDigit(rune(s[i])) {
		tmp := string(s[i])
		i++
		for i < len(s) && unicode.IsDigit(rune(s[i])) {
			tmp += string(s[i])
			i++
		}
		a, _ := strconv.Atoi(tmp)
		root.right = &Node{a, depth + 1, nil, nil}
	} else {
		j := findMatchingIndex(s, i, len(s)-1)
		root.right = TreeFromString(s[i+1:j], depth+1)
	}
	return root
}

func findMatchingIndex(s string, si int, ei int) int {
	if si > ei {
		return -1
	}
	var stack []rune
	for i := si; i <= ei; i++ {
		if s[i] == '[' {
			stack = append(stack, '[')
		} else if s[i] == ']' {
			if stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
				if len(stack) == 0 {
					return i
				}
			}
		}
	}
	return -1
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
	v := traverse(n, nil, nil)
	for v {
		v = traverse(n, nil, nil)
	}
}

func magnitude(node *Node) int {
	if node.value != -1 {
		return node.value
	}
	return 3*magnitude(node.left) + 2*magnitude(node.right)
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

func get_right_nb(n *Node) *Node {
	if n.value != -1 {
		return n
	}
	if n.left != nil {
		return get_right_nb(n.left)
	}
	if n.right != nil {
		return get_right_nb(n.right)
	}
	return nil
}

func get_left_nb(n *Node) *Node {
	if n.value != -1 {
		return n
	}
	if n.right != nil {
		return get_left_nb(n.right)
	}
	if n.left != nil {
		return get_left_nb(n.left)
	}
	return nil
}

func traverse(n *Node, left_nb *Node, right_nb *Node) bool {
	if n.left != nil {
		s := traverse(n.left, left_nb, n.right)
		if s {
			return true
		}
	}

	if n.depth > 3 && n.value == -1 &&
		n.left != nil && n.left.value != -1 &&
		n.right != nil && n.right.value != -1 {
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
		if n.value%2 == 1 {
			rv++
		}
		l := &Node{lv, n.depth + 1, nil, nil}
		r := &Node{rv, n.depth + 1, nil, nil}
		*n = Node{-1, n.depth, l, r}
		return true
	}

	if n.right != nil {
		s := traverse(n.right, n.left, right_nb)
		if s {
			return true
		}
	}
	return false
}
func LoadData(filename string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}
	return strings.Split(strings.Trim(string(fileBytes),"\n"), "\n") , nil

}

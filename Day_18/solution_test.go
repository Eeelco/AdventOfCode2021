package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomework(t *testing.T) {
    assert := assert.New(t)
    data, err := LoadData("test_data_0.txt")
    assert.Equal(err, nil)
    d := AddNumbers(data)
    assert.Equal(print_tree(d), "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")


    data, err = LoadData("test_data_1.txt")
    assert.Equal(err, nil)
    d = AddNumbers(data)
    assert.Equal(print_tree(d), "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]")
    // data, err = LoadData("test_data.txt")
    // assert.Equal(err, nil)
    // d = AddNumbers(data)
    // assert.Equal(print_tree(d), "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]")
}

func TestAddition(t *testing.T) {
    assert := assert.New(t)
    g := []string{"[1,1]", 
"[2,2]",
"[3,3]",
"[4,4]",
}
assert.Equal(print_tree(AddNumbers(g)), "[[[[1,1],[2,2]],[3,3]],[4,4]]")
h := []string{"[1,1]", 
"[2,2]",
"[3,3]",
"[4,4]",
"[5,5]",
"[6,6]",
}
assert.Equal(print_tree(AddNumbers(h)), "[[[[5,0],[7,4]],[5,5]],[6,6]]")

a := MakeTree("[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]")
b := MakeTree("[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]")
c := add(a,b)
reduce(c)
assert.Equal(print_tree(c), "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]")

}

func TestSplit(t *testing.T) {
    assert := assert.New(t)
    c := MakeTree("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]")
    traverse(c,nil,nil)
    assert.Equal(print_tree(c),"[[3,[2,[8,0]]],[9,[5,[7,0]]]]")
    c = MakeTree("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]")
    traverse(c,nil,nil)
    assert.Equal(print_tree(c),"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]")
    c = MakeTree("[7,[6,[5,[4,[3,2]]]]]")
    traverse(c,nil,nil)
    assert.Equal(print_tree(c),"[7,[6,[5,[7,0]]]]")
    c = MakeTree("[[[[[9,8],1],2],3],4]")
    traverse(c,nil,nil)
    assert.Equal(print_tree(c),"[[[[0,9],2],3],4]")
}

func TestTreefuncs(t *testing.T) {
    assert := assert.New(t)
    a := MakeTree("[[[[4,3],4],4],[7,[[8,4],9]]]")
    b := MakeTree("[1,1]")
    c := add(a, b)
    assert.Equal(print_tree(c),"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]")
    traverse(c,nil,nil)
    assert.Equal(print_tree(c), "[[[[0,7],4],[7,[[8,4],9]]],[1,1]]")
    traverse(c,nil,nil)
    assert.Equal(print_tree(c), "[[[[0,7],4],[15,[0,13]]],[1,1]]")
    traverse(c,nil,nil)
    assert.Equal(print_tree(c), "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]")
    traverse(c,nil,nil)
    assert.Equal(print_tree(c), "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]")
    traverse(c,nil,nil)
    assert.Equal(print_tree(c), "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")
}

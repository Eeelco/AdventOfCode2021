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
    data, err = LoadData("test_data.txt")
    assert.Equal(err, nil)
    d = AddNumbers(data)
    assert.Equal(print_tree(d), "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]")
}


package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
    cubes, _ := LoadData("input.txt")
    b := Box{[]Cube{}}
    b.Fold(cubes)
    fmt.Printf("Part 2 solution:%d\n\n", b.Volume())
}


type Site struct{
    x int
    y int
    z int
}

type Box struct {
    cubes []Cube
}

func (b *Box) Fold(cubes []Cube) {
    for _, c := range cubes {
        b.Toggle(c)
    }
}

func (b *Box) Volume() uint64 {
    vol := uint64(0)
    for _, c := range b.cubes {
        vol += uint64(c.Volume())
    }
    return vol
}


func (b *Box) Toggle(c1 Cube) {
    var tmp_cubes []Cube
    for _, c2 := range b.cubes {
        if c1.Overlaps(c2) {
            tmp_cubes = append(tmp_cubes, c2.Split(c1)...)
        } else {
            tmp_cubes = append(tmp_cubes, c2)
        }
    }
    if c1.action == "on" {
        tmp_cubes = append(tmp_cubes, c1)
    }
    b.cubes = tmp_cubes
}


type Cube struct {
    action string
    xmin float64
    xmax float64
    ymin float64
    ymax float64
    zmin float64
    zmax float64
}

func (c1 *Cube)Split(c2 Cube) []Cube {
    var new_cubes []Cube
    if c1.xmin < c2.xmin {
        new_cubes = append(new_cubes, Cube{"on",c1.xmin, c2.xmin-1, c1.ymin, c1.ymax, c1.zmin, c1.zmax})
    }
    if c1.xmax > c2.xmax {
        new_cubes = append(new_cubes, Cube{"on",c2.xmax+1, c1.xmax, c1.ymin, c1.ymax, c1.zmin, c1.zmax})
    }
    if c1.ymin < c2.ymin {
        new_cubes = append(new_cubes, Cube{"on",math.Max(c1.xmin, c2.xmin), math.Min(c1.xmax, c2.xmax), c1.ymin, c2.ymin-1, c1.zmin, c1.zmax})
    }
    if c1.ymax > c2.ymax {
        new_cubes = append(new_cubes, Cube{"on",math.Max(c1.xmin, c2.xmin), math.Min(c1.xmax, c2.xmax), c2.ymax+1, c1.ymax, c1.zmin, c1.zmax})
    }
    if c1.zmin < c2.zmin {
        new_cubes = append(new_cubes, Cube{"on",math.Max(c1.xmin, c2.xmin), math.Min(c1.xmax, c2.xmax), math.Max(c1.ymin, c2.ymin), math.Min(c1.ymax, c2.ymax), c1.zmin, c2.zmin-1})
    }
    if c1.zmax > c2.zmax {
        new_cubes = append(new_cubes, Cube{"on",math.Max(c1.xmin, c2.xmin), math.Min(c1.xmax, c2.xmax), math.Max(c1.ymin, c2.ymin), math.Min(c1.ymax, c2.ymax), c2.zmax+1, c1.zmax})
    }
    return new_cubes
}

func (c1 *Cube) Intersection(c2 Cube) Cube {
    return Cube{"", math.Max(c1.xmin, c2.xmin), math.Min(c1.xmax, c2.xmax), math.Max(c1.ymin, c2.ymin), math.Min(c1.ymax, c2.ymax), math.Max(c1.zmin, c2.zmin), math.Min(c1.zmax, c2.zmax)}
}

func (c1 *Cube) Overlaps(c2 Cube) bool {
    return c2.xmin <= c1.xmax && c2.ymin <= c1.ymax && c2.zmin <= c1.zmax && c2.xmax >= c1.xmin && c2.ymax >= c1.ymin && c2.zmax >=c1.zmin
}

func (c *Cube) Volume() uint64 {
    return uint64((c.xmax + 1.0 - c.xmin)*(c.ymax + 1.0 - c.ymin)*(c.zmax + 1.0 - c.zmin))
}


func LoadData(filename string) ([]Cube, error) {
    fileBytes, err := ioutil.ReadFile(filename)

    if err != nil {
        return nil, err
    }
    lines :=  strings.Split(strings.Trim(string(fileBytes),"\n"), "\n")
    var cubes []Cube
    for _, l := range lines {
        tmp := strings.Split(l, " ")
        ranges := strings.Split(tmp[1], ",")
        var vals []float64
        for _, r := range ranges {
            a := strings.Split(r,"=")[1]
            b := strings.Split(a, "..")
            min, _ := strconv.ParseFloat(b[0], 64)
            max, _ := strconv.ParseFloat(b[1], 64)
            vals = append(vals, min)
            vals = append(vals, max)
        }
        cubes = append(cubes, Cube{tmp[0], vals[0], vals[1], vals[2], vals[3], vals[4], vals[5]})
    }
    return cubes, nil
}



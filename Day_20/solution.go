package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Image [][]rune

func main() {
    enhancement, image, err := LoadData("test_data.txt")
    input_enhancement, input_image, _ := LoadData("input.txt")
    if err != nil {
        fmt.Println("Error opening test file")
        os.Exit(1)
    }
    for i := 0; i < 2; i++ {
        image = NextStep(image, enhancement, i)
        input_image = NextStep(input_image, input_enhancement, i)
    }
    lights := CountLit(image)
    input_lights := CountLit(input_image)
    if lights != 35 {
        fmt.Printf("Test data error. Expected 35, got %d", lights)
    }
    fmt.Printf("Part 1 solution:\n\n%d\n\n", input_lights)

    for i := 0; i < 48; i++ {
        image = NextStep(image, enhancement, i)
        input_image = NextStep(input_image, input_enhancement, i)
    }
    lights = CountLit(image)
    input_lights = CountLit(input_image)
    if lights != 3351 {
        fmt.Printf("Test data error. Expected 3351, got %d", lights)
    }
    fmt.Printf("Part 2 solution:\n\n%d\n\n", input_lights)
}

func NextStep(img Image, enhancement string, step_nr int) Image {
    padChar := '.'
    if step_nr%2 == 1 && enhancement[0] == '#' {
        padChar = '#'
    }
    padded := PadImage(img, 2, padChar)
    next := EmptyImage(len(padded)-2, len(padded[0])-2)

    for i := 1; i < len(padded) - 1; i++ {
        for j := 1; j < len(padded[0])-1; j++ {
            val := CalcPixel(&padded, i, j)
            next[i-1][j-1] = rune(enhancement[val])
        }
    }
    return next
}

func EmptyImage(ydims int, xdims int) Image {
    out := make(Image, ydims) 
    for i := 0; i < ydims; i++ {
        out[i] = make([]rune, xdims)
        for j := 0; j < xdims; j++ {
            out[i][j] = '.'
        }
    }
    return out
}

func PadImage(img Image, n int, padChar rune) Image {
    output := make(Image, len(img) + 2*n)
    for i := 0; i < len(img) + 2*n; i++ {
        output[i] = make([]rune, len(img[0]) + 2*n)
        for j := 0; j < len(img[0]) + 2*n; j++ {
            if i < n || j < n || i >= len(img) + n || j >= len(img[0]) + n {
                output[i][j] = padChar
            } else {
                output[i][j] = img[i-n][j-n]
            }
        }
    }
    return output
}

func CalcPixel(img *Image, centerx int, centery int) int64{
    out := ""
    for i := centerx - 1; i < centerx + 2; i ++ {
        for j := centery - 1; j < centery + 2; j++ {
            if (*img)[i][j] == '.' {
                out += string('0')
            } else {
                out += string('1')
            }
        }
    }
    nr, _ := strconv.ParseInt(out, 2, 64)
    return nr
}

func CountLit(img Image) int {
    o := 0
    for i := 0; i < len(img); i++ {
        for j := 0; j < len(img[0]); j++ {
            if img[i][j] == '#' {
                o++
            }
        }
    }
    return o
}

func LoadData(filename string) (string, Image, error) {
    fileBytes, err := ioutil.ReadFile(filename)

    if err != nil {
        return "", nil, err
    }
    lines :=  strings.Split(strings.Trim(string(fileBytes),"\n"), "\n")
    enhancement := lines[0]
    outfile := make(Image, len(lines) - 2)
    for i := 0; i < len(lines)-2; i++ {
        outfile[i] = make([]rune, len(lines[2]))
        for j := 0; j < len(lines[2]); j++ {
            outfile[i][j] = rune(lines[i+2][j])
        }
    } 
    return enhancement, outfile, nil
}

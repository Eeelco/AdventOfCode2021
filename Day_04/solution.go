package day04

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func In(lst []int, val int) bool {
    for _,j := range lst {
        if j == val {
            return true
        }
    }
    return false
}

func FindLoser(moves []int, boards [][][]int) int {
    board_size := len(boards[0])
    found := make([][][]bool,len(boards))
    var winners []int
    var winner_scores []int
    for i := 0; i < len(boards); i++ {
        found[i] = make([][]bool, board_size)
        for j := 0; j < board_size; j++ {
            found[i][j] = make([]bool, board_size)
        }
    }
    for _, m := range moves {
        for i, b := range boards {
            x, y := FindIndex(b, m)
            if x >= 0 {
                found[i][x][y] = true
            }
            if CheckBingo(found[i]) && !In(winners, i){
                winners = append(winners, i)
                winner_scores = append(winner_scores,CalcScore(b, found[i], m))
            }
        }
    }
    return winner_scores[len(winner_scores)-1]
}

func FindWinner(moves []int, boards [][][]int) int {
    board_size := len(boards[0])
    found := make([][][]bool,len(boards))
    for i := 0; i < len(boards); i++ {
        found[i] = make([][]bool, board_size)
        for j := 0; j < board_size; j++ {
            found[i][j] = make([]bool, board_size)
        }
    }
    for _, m := range moves {
        for i, b := range boards {
            x, y := FindIndex(b, m)
            if x >= 0 {
                found[i][x][y] = true
            }
            if CheckBingo(found[i]) {
                return CalcScore(b, found[i], m)
            }
        }
    }
    return -1
}

func LoadData(filename string) ([]int, [][][]int, error) {
    fileBytes, err := ioutil.ReadFile(filename)

    if err != nil {
        return nil, nil, err
    }
    raw_data := strings.Split(strings.Trim(string(fileBytes),"\n"), "\n")
    moves := strings.Split(raw_data[0], ",")
    Intmoves := make([]int, len(moves))
    for i, v := range moves {
        Intmoves[i],_ = strconv.Atoi(v)
    }
    var boards [][][]int
    board_size := len(strings.Fields(raw_data[2]))
    for i := 1; i < len(raw_data); i++ {
        if len(raw_data[i]) == 0 && i < len(raw_data)-1 {
            boards = append(boards, make([][]int, 0))
        } else {
            tmp := strings.Fields(raw_data[i])
            boards[len(boards)-1] = append(boards[len(boards)-1], make([]int, board_size))
            blen := len(boards[len(boards)-1]) - 1
            for i,v := range tmp {
                boards[len(boards)-1][blen][i],_ = strconv.Atoi(v)
            }
        }
    }
    return Intmoves,boards,nil
}

func FindIndex(board [][]int, item int) (int, int) {
    x, y := -1, -1
    for i, v := range board {
        for j, w := range v {
            if w == item {
                x = i
                y = j
            }
        }
    }
    return x, y
}

func CheckBingo(oks [][]bool) bool {
    blen := len(oks)

    for i := 0; i < blen; i++ {
        bcount := 0
        for j := 0; j < blen; j++ {
            if oks[i][j] {
                bcount++
            }
        }
        if bcount == blen {
            return true
        }
    }
    for i := 0; i < blen; i++ {
        bcount := 0
        for j := 0; j < blen; j++ {
            if oks[j][i] {
                bcount++
            }
        }
        if bcount == blen {
            return true
        }
    }
    return false
}

func CalcScore(board [][]int, found [][]bool, last_number int) int {
    sum := 0
    for i, v := range board {
        for j, w := range v {
            if !found[i][j] {
                sum += w
            }
        }
    }
    return last_number * sum
}

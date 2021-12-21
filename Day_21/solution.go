package main

import (
	"fmt"
	"sync"
)

type State struct {
    player1pos int
    player2pos int
    player1score int
    player2score int
    n_turns int
    nr_ways int
}
type Result struct {
    p1wins int
    p2wins int
}
type Roll struct {
    value int
    nr_ways int
}

func main() {
    test_result := play_game(4,8)
    fmt.Printf("Part 1 test_result:\n\n%d\n\n", test_result)
    test_result = play_game(6,10)
    fmt.Printf("Part 1 solution:\n\n%d\n\n", test_result)

    test2_result := parallel_quantum_game(4,8)
    fmt.Printf("Part 2 test_result:\n\n%d\n\n", test2_result)
    test2_result = parallel_quantum_game(6,10)
    fmt.Printf("Part 2 solution:\n\n%d\n\n", test2_result)
}

func parallel_quantum_game(player1pos int, player2pos int) int {
    possible_rolls := rolls()
    ch := make(chan Result, 7)
    ini_state := State{player1pos, player2pos, 0,0,0 , 1}
    var wg sync.WaitGroup
    wg.Add(7)
    for _, r := range possible_rolls {
        next, _ := advance(ini_state, r, nil)
        go quantum_game(next, &wg, ch)
    }
    wg.Wait()
    close(ch)
    p1wins := 0
    p2wins := 0
    for v := range ch {
        p1wins += v.p1wins
        p2wins += v.p2wins
    }
    if p1wins > p2wins {
        return p1wins
    } else {
        return p2wins
    }
}


func quantum_game(init_state State, wg *sync.WaitGroup, ch chan Result) {
    defer wg.Done()
    states := []State{init_state}
    wins := make(map[int]int)
    wins[1] = 0
    wins[2] = 0
    possible_rolls := rolls()

    for len(states) > 0 {
        s := states[len(states)-1]
        states = states[:len(states)-1]
        for _, r := range possible_rolls {
            next_s, done := advance(s, r, &wins)
            if !done {
                states = append(states, next_s)
            }
        }
    }
    ch <- Result{wins[1], wins[2]}
}

func advance(state State, roll Roll, wins *map[int]int) (State, bool) {
    var score, pos *int
    if state.n_turns % 2 == 0 {
        score, pos = &state.player1score, &state.player1pos
    } else {
        score, pos = &state.player2score, &state.player2pos
    }
    *pos += roll.value
    for *pos > 10 {
        *pos -= 10
    }
    *score += *pos
    state.n_turns += 3
    state.nr_ways *= roll.nr_ways
    if state.player1score >= 21 {
        (*wins)[1] += state.nr_ways
        return State{}, true
    } else if state.player2score >= 21 {
        (*wins)[2]  += state.nr_ways
        return State{}, true
    }
    return state, false
}

func play_game(player1pos int, player2pos int) int {
    player1score, player2score, die_state, n_turns := 0, 0, 1, 0
    for player1score < 1000 && player2score < 1000 {
        var score, pos *int
        if n_turns % 2 == 0 {
            score, pos = &player1score, &player1pos
        } else {
            score, pos = &player2score, &player2pos
        }
        c := 0
        for i := 0; i < 3; i++ {
            c += die_state
            die_state++
            if die_state > 100 {
                die_state = 1
            }
        }
        *pos += c
        for *pos > 10 {
            *pos -= 10
        }
        *score += *pos
        n_turns += 3
    }
    loser := player1score
    if player2score < player1score {
        loser = player2score
    }
    return loser * n_turns
}

func rolls() []Roll {
    return []Roll{
        Roll{3, 1},
        Roll{4, 3},
        Roll{5, 6},
        Roll{6, 7},
        Roll{7, 6},
        Roll{8, 3},
        Roll{9, 1},
    }
}

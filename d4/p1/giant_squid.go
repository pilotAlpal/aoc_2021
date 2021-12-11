package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
//	"regexp"
//    "errors"
    "strings"
    "strconv"
)

type Cell struct {
    N int
    Ticked bool
}

type Board struct {
    Cells [][]Cell
    Values []int
}

func parse_cell(cell_n int) Cell {
    return Cell {cell_n, false}
}

func parse_row(row, sep string) ([]Cell, []int) {
    splitted := strings.Split(strings.TrimPrefix(strings.ReplaceAll(row, "  ", " "), " "), sep)
    var out_row []Cell
    var row_values []int
    for _, v := range splitted {
        if s, err := strconv.Atoi(v); err == nil {
            out_row = append(out_row, parse_cell(s))
            row_values = append(row_values, s)
	    } else {
            log.Fatal(err)
        }
    }
    return out_row, row_values
}

func parse_sequence(in_seq string) []int {
    splitted := strings.Split(strings.TrimPrefix(strings.ReplaceAll(in_seq, "  ", " "), " "), ",")
    var out_sq []int
    for _, v := range splitted {
        if s, err := strconv.Atoi(v); err == nil {
            out_sq = append(out_sq, s)
	    } else {
            log.Fatal(err)
        }
    }
    return out_sq
}

func parse_board(scanner *bufio.Scanner) Board {
    rows := make([][]Cell, 5)
    var b_values []int
    for i, _ := range rows {
        scanner.Scan();
        row, r_values := parse_row(scanner.Text(), " ")
        rows[i] = row
        b_values = append(b_values, r_values...)
    }
    return Board{rows, b_values}
}
func parse_input(input_path string) ([]int, []Board) {
    file, err := os.Open(input_path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan();
    seq := parse_sequence(scanner.Text())
    var boards []Board
    for ;scanner.Scan(); {
        boards = append(boards, parse_board(scanner))
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return seq, boards
}

func (b Board) fill_row(r int) bool {
    for _, c := range b.Cells[r] {
       if !c.Ticked {
        return false
       }
    }
    return true
}

func (b Board) fill_column(c int) bool {
    for _, r := range b.Cells {
       if !r[c].Ticked {
        return false
       }
    }
    return true
}

func (b Board) tickNumber(n int) (Board, bool) {
    n_in_b := false
    for _, v := range b.Values {
       if n == v {
            n_in_b = true
       }
    }
    if n_in_b{
        for i, r := range b.Cells {
            for j, c := range r {
                if n == c.N {
                    b.Cells[i][j] = Cell{n, true}
                    if b.fill_row(i) || b.fill_column(j) {
                        return b, true
                    }
                }
            }
        }
    }
    return b, false
}

func tick_boards(seq []int, boards []Board) (Board, int){
    for _, v := range seq {
        for _, b := range boards {
            b, b_wins := b.tickNumber(v)
            if b_wins {
                return b, v
            }
        }
    }
    return Board{}, -1

}

func (b Board) sum_unmarked() int {
    s := 0
    for _, r := range b.Cells {
        for _, c := range r {
            if !c.Ticked {
                s = s + c.N
            }
        }
    }
    return s
}

func main() {
    sequence, boards := parse_input("../inputs/input.in")
    winner_b, last_call := tick_boards(sequence, boards)
    fmt.Println(winner_b.sum_unmarked() * last_call)
}


package main

import (
    "fmt"
    "regexp"
    "os"
    "log"
    "bufio"
    "strconv"
    "math"
)

func line_matches_regexp(in_l string) bool {
    expected_re := `[0-9]{10}`
    re := regexp.MustCompile(expected_re)
    return re.Match([]byte(in_l))
}

func get_row_values(line string) []int {
    var r_values []int
    for _, c := range line {
        if s, err := strconv.Atoi(string(c)); err == nil {
            r_values = append(r_values, s)
        } else {
            log.Fatal(c, "Could not convert to int")
            log.Fatal(err)
        }
    }
    return r_values
}

func parse_input(input_path string) [][]int {
    file, err := os.Open(input_path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)

    var e_levels [][]int
    for scanner.Scan() {
        l := scanner.Text()
        if line_matches_regexp(l) {
            e_levels = append(e_levels, get_row_values(l))
        } else {
            log.Fatal(l, "Didn't match expected regex")
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return e_levels
}

func get_adjacents(i, j, size int)  ([]int, []int) {
    var adj_i, adj_j []int
    incrs := [8][2]int{ {-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1} }
    for _, inc := range incrs {
        i_pos := i + inc[0]
        j_pos := j + inc[1]
        if 0 <= i_pos && i_pos < size && 0 <= j_pos && j_pos < size {
            adj_i = append(adj_i, i_pos)
            adj_j = append(adj_j, j_pos)
        }
    }
    return adj_i, adj_j
}

type Position struct {
    I,J int
}

func take_steps(e_levels [][]int, steps int) ([][]int, int) {
    size := len(e_levels)
    for x := 0; x < steps; x++ {
        var to_be_flashed []Position
        xth_flashes := make(map[Position]bool)
        for i := 0; i < size; i++ {
            for j := 0; j < size; j++ {
                e_levels[i][j]++
                if e_levels[i][j] > 9 {
                    to_be_flashed = append(to_be_flashed, Position{i, j})
                    xth_flashes[Position{i, j}] = true
                }
            }
        }
        for len(to_be_flashed) > 0 {
            i, j := to_be_flashed[0].I, to_be_flashed[0].J
            to_be_flashed = to_be_flashed[1:]
            adjs_i, adjs_j := get_adjacents(i, j, size)
            for n, v_i := range adjs_i {
                if _, included := xth_flashes[Position{v_i, adjs_j[n]}]; !included {
                    e_levels[v_i][adjs_j[n]]++
                    if e_levels[v_i][adjs_j[n]] > 9 {
                        new_flash := Position{v_i, adjs_j[n]}
                        to_be_flashed = append(to_be_flashed, new_flash)
                        xth_flashes[new_flash] = true
                    }
                }
            }
            e_levels[i][j] = 0
        }
        if len(xth_flashes) == 100 {
            return e_levels, x + 1
        }
    }
    return e_levels, -1
}

func main() {
    energy_levels := parse_input("../inputs/input.in")
    all_0s_step := 0
    _, all_0s_step = take_steps(energy_levels, math.MaxInt64)
    fmt.Println(all_0s_step)
}


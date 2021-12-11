package main

import (
    "fmt"
    "os"
    "bufio"
    "regexp"
    "log"
    "strings"
    "strconv"
)

func line_matches_regexp(in_l string) bool {
    expected_re := `[0-9]+`
    re := regexp.MustCompile(expected_re)
    return re.Match([]byte(in_l))
}

func as_ints(str_sl []string) []int {
    var ints_sl []int
    for _, v := range str_sl {
        if s, err := strconv.Atoi(v); err == nil {
            ints_sl = append(ints_sl, s)

        } else {
             log.Fatal(err)
             log.Fatal(v, "Couldn't parse as int")
        }

    }
    return ints_sl
}


func parse_input(input_path string) [][]int{
    file, err := os.Open(input_path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    var s_flows [][]int
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        l := scanner.Text()
        if line_matches_regexp(l) {
            s_flows = append(s_flows, as_ints(strings.Split(l, "")))
        } else {
            log.Fatal(l, "Didn't match expected regex")
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return s_flows
}

func is_low_point(smoke_flows [][]int, i,j int) bool {
    if i > 0 && smoke_flows[i - 1][j] <= smoke_flows[i][j] {
        return false
    } else if j > 0 && smoke_flows[i][j - 1] <= smoke_flows[i][j] {
        return false
    } else if i < len(smoke_flows) -1 && smoke_flows[i + 1][j] <= smoke_flows[i][j] {
        return false
    } else if j < len(smoke_flows[i]) -1 && smoke_flows[i][j + 1] <= smoke_flows[i][j] {
        return false
    }
    return true
}

func get_low_points(s_flows [][]int) ([]int, []int, int) {
    var l_points_i, l_points_j []int
    risk_sum := 0
    for i, v := range s_flows {
        for j, _ := range v {
            if is_low_point(s_flows, i, j) {
                l_points_i = append(l_points_i, i)
                l_points_j = append(l_points_j, j)
                risk_sum += 1 + s_flows[i][j]
            }
        }
    }
    return l_points_i, l_points_j, risk_sum
}

func main() {
    smoke_flows := parse_input("../inputs/input.in")
    _, _, risk_sum := get_low_points(smoke_flows)
    fmt.Println(risk_sum)
}


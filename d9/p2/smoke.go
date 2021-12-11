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

type Point struct {
    I,J int
}

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

func contains_point(points []Point, i,j int) bool {
    for _, v := range points {
        if v.I == i && v.J == j {
            return true
        }
    }
    return false
}

func get_basin(s_flows [][]int, i,j int) []Point {
    var basin []Point
    basin = append(basin, Point{i, j})
    for x:=0; x < len(basin); x++ {
        ns := get_neighbors(s_flows, basin[x].I, basin[x].J)
        for _, p := range ns {
            if s_flows[p.I][p.J] < 9 && !contains_point(basin, p.I, p.J) {
                basin = append(basin, p)
            }
        }
    }
    return basin
}

func get_neighbors(smoke_flows [][]int, i,j int) []Point {
    var neighs []Point
    if i > 0 {
        neighs = append(neighs, Point{i - 1, j})
    }
    if j > 0 {
        neighs = append(neighs, Point{i, j - 1})
    }
    if i < len(smoke_flows) -1 {
        neighs = append(neighs, Point{i + 1, j})
    }
    if j < len(smoke_flows[i]) -1 {
        neighs = append(neighs, Point{i, j + 1})
    }
    return neighs

}

func get_low_points(s_flows [][]int) ([]Point, [][]Point) {
    var low_points []Point
    var basins [][]Point
    for i, v := range s_flows {
        for j, _ := range v {
            if is_low_point(s_flows, i, j) {
                low_points = append(low_points, Point{i, j})
                basins = append(basins, get_basin(s_flows, i, j))
            }
        }
    }
    return low_points, basins
}

func get_3_longest(basins [][]Point) ([]Point, []Point, []Point) {
    var x, y, z []Point
    for _, b := range basins {
        lb := len(b)
        if lb > len(x) {
            z = y
            y = x
            x = b
        } else if lb > len(y) {
            z = y
            y = b
        } else if lb > len(z) {
            z = b
        }
    }
    return x, y, z

}

func main() {
    smoke_flows := parse_input("../inputs/input.in")
    _, basins := get_low_points(smoke_flows)

    a, b, c := get_3_longest(basins)

    fmt.Println(len(a) * len(b) * len(c))
}


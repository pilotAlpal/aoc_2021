package main

import (
    "fmt"
    "regexp"
    "os"
    "log"
    "strconv"
    "bufio"
    "strings"
)

const POINTS_RE = `([0-9]+),([0-9]+)`
const FOLDS_RE = `fold along ([x|y])=([0-9]+)`

type Point struct {
    X,Y int
}

func line_matches_regexp(in_l, expected_re string) bool {
    re := regexp.MustCompile(expected_re)
    return re.Match([]byte(in_l))
}

func get_point(point, expected_re string) Point {
    x, y := -1, -1
    re := regexp.MustCompile(expected_re)
    x_y := strings.Split(re.FindStringSubmatch(point)[0], ",")
    if sy, err := strconv.Atoi(string(x_y[1])); err == nil {
        y = sy
    } else {
        log.Fatal(err)
    }
    if sx, err := strconv.Atoi(string(x_y[0])); err == nil {
        x = sx
    } else {
        log.Fatal(err)
    }

    return Point{x, y}
}

func get_fold(fold, expected_re string) Point {
    re := regexp.MustCompile(expected_re)
    f := re.FindStringSubmatch(fold)[1:]
    v := 0
    if s, err := strconv.Atoi(string(f[1])); err == nil {
        v = s
    } else {
        log.Fatal(err)
    }
    if string(f[0]) == "x" {
        return Point{v, 0}
    }
    return Point{0, v}
}


func parse_input(input_path string) (map[Point]bool, []Point) {
    file, err := os.Open(input_path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)

    points := make(map[Point]bool)
    read_folds := false
    for !read_folds {
        scanner.Scan()
        l := scanner.Text()
        if line_matches_regexp(l, POINTS_RE) {
            points[get_point(l, POINTS_RE)] = true
        } else if l == "" {
            read_folds = true
        } else {
            log.Fatal(l, "Didn't match expected regex", POINTS_RE)
        }
    }
    var folds []Point
    for scanner.Scan()  {
        l := scanner.Text()
        if line_matches_regexp(l, FOLDS_RE) {
            folds = append(folds, get_fold(l, FOLDS_RE))
        } else {
            log.Fatal(l, "Didn't match expected regex", FOLDS_RE)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return points, folds
}

func take_fold(positions map[Point]bool, fold Point) map[Point]bool {
    n_positions := make(map[Point]bool)
    x, y := fold.X, fold.Y
    if x > 0 {
        for p, _ := range positions {
            if p.X < x {
                n_positions[p] = true
            } else {
                n_positions[Point{2 * x - p.X, p.Y}] = true
            }
        }
    } else if y > 0 {
        for p, _ := range positions {
            if p.Y < y {
                n_positions[p] = true
            } else {
                n_positions[Point{p.X, 2 * y - p.Y}] = true
            }
        }
    }
    return n_positions
}

func main() {
//    points, folds := parse_input("../inputs/input.ex")
    points, folds := parse_input("../inputs/input.in")
//    fmt.Println(points)
    one_fold := take_fold(points, folds[0])
    fmt.Println(len(one_fold))
}


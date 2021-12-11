package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
	"regexp"
    "math"
    "strings"
    "strconv"
)

type Board struct {
    Cells [][]int
    MoreThanTwo int
}

func get_coordinates(point string) (int, int) {
    x := -1
    y := -1
    values := strings.Split(point, ",")
    if s, err := strconv.Atoi(values[0]); err == nil {
        x = s
    } else {
        log.Fatal(err)
        log.Fatal(values[0], "Couldn't parse as int")
    }
    if s, err := strconv.Atoi(values[1]); err == nil {
        y = s
    } else {
        log.Fatal(err)
        log.Fatal(values[1], "Couldn't parse as int")
    }
    return x, y
}

func get_max_min(a, b int) (int, int) {
    if a < b {
        return b, a
    }
    return a, b
}

func get_incr(a, b int) int {
    if a < b {
        return 1
    }
    return -1
}


func (b Board) increase_point_count(x, y int) Board {
    for i := len(b.Cells); i <= y; i++ {
        b.Cells = append(b.Cells, make([]int, 1))
    }
    if x >= len(b.Cells[y]) {
        for i := len(b.Cells[y]); i < x; i++ {
            b.Cells[y] = append(b.Cells[y], 0)
        }
        b.Cells[y] = append(b.Cells[y], 1)
    } else {
        b.Cells[y][x] = b.Cells[y][x] + 1
        if b.Cells[y][x] == 2 {
            b.MoreThanTwo++
        }
    }
    return b

}

func (b Board) include_diagonal(x0, y0, x1, y1 int) Board {
    x_incr := get_incr(x0, x1)
    y_incr := get_incr(y0, y1)

    for i := 0; float64(i) <= math.Abs(float64(y0 - y1)); i++ {
        b = b.increase_point_count(x0 + (i * x_incr), y0 + (i * y_incr))
    }
    return b
}

func (b Board) include_horizontal(y, x0, x1 int) Board {
    h, l := get_max_min(x0, x1)
    for i := l; i <= h; i++ {
        b = b.increase_point_count(i, y)
    }
    return b
}

func (b Board) include_vertical(x, y0, y1 int) Board {
    h, l := get_max_min(y0, y1)
    for i := l; i <= h; i++ {
        b = b.increase_point_count(x, i)
    }
    return b
}

func (b Board) include_points(x0, y0, x1, y1 int) Board {
    if x0 == x1 {
        b = b.include_vertical(x0, y0, y1)
    } else if y0 == y1 {
        b = b.include_horizontal(y0, x0, x1)
    } else {
        b = b.include_diagonal(x0, y0, x1, y1)
    }
    return b
}

func (b Board) add_line(line string) Board{
    segment := strings.Split(line, " -> ")
    p0x, p0y  := get_coordinates(segment[0])
    p1x, p1y  := get_coordinates(segment[1])
    return b.include_points(p0x, p0y, p1x, p1y)
}

func matches_regexp(input string) bool {
    re := regexp.MustCompile(`[0-9]{1,3},[0-9]{1,3} -> [0-9]{1,3},[0-9]{1,3}`)
    return re.Match([]byte(input))
}

func parse_input(input_path string) Board {
    file, err := os.Open(input_path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    board := Board{}
    for ;scanner.Scan(); {
        if matches_regexp(scanner.Text()){
            board = board.add_line(scanner.Text())
        } else {
            fmt.Println(scanner.Text())
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return board
}


func main() {
    board := parse_input("../inputs/input.in")
    fmt.Println(board.MoreThanTwo)
}


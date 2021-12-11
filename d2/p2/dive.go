package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
	"regexp"
    "strings"
    "strconv"
)

type Coordinate struct {
	Depth int
	HorizontalPosition int
    Aim int
}

func get_position(c0, c1 Coordinate) Coordinate {
    return Coordinate{c0.Depth + c1.Depth, c0.HorizontalPosition + c1.HorizontalPosition, c0.Aim + c1.Aim}
}


func multiply_indexes(move Coordinate) int {
    return move.Depth * move.HorizontalPosition
}

func parse_command(command string, current_aim int) Coordinate {
	re := regexp.MustCompile(`[forward|down|up] [1-9]`)
    if re.Match([]byte(command)){
        args := strings.Split(command, " ")
        dir, amount := args[0], args[1]
        if s, err := strconv.Atoi(amount); err == nil {
            if dir == "forward" {
                return Coordinate{current_aim * s, s, 0}
            } else if dir == "up" {
                return Coordinate{0, 0, -s}
            } else {
                return Coordinate{0, 0, s}
            }

        } else {
            log.Fatal(err)
            log.Fatal(amount, "Couldn't parse as int")
        }
    } else {
        log.Fatal(command, "Didn't match expected regex")
    }
    return Coordinate{0, 0, 0}
}

func parse_input(input_path string) Coordinate{
    file, err := os.Open(input_path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    current_pos := Coordinate{0, 0, 0}
    for scanner.Scan() {
        current_pos = get_position(current_pos, parse_command(scanner.Text(), current_pos.Aim))
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return current_pos

}

func main() {
    position := parse_input("../inputs/input.in")
    fmt.Println(multiply_indexes(position))

}

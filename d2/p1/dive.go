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

type Movement struct {
	X int
	Y int
}

func get_position(moves []Movement) Movement {
    current_pos := Movement{0, 0}
    for _, m := range moves {
        current_pos = Movement{current_pos.X + m.X, current_pos.Y + m.Y}
    }
    return current_pos
}


func multiply_indexes(move Movement) int {
    return move.X * move.Y
}

func parse_command(command string) Movement {
	re := regexp.MustCompile(`[forward|down|up] [1-9]`)
    if re.Match([]byte(command)){
        args := strings.Split(command, " ")
        dir, amount := args[0], args[1]
        if s, err := strconv.Atoi(amount); err == nil {
            if dir == "forward" {
                return Movement{0, s}
            } else if dir == "up" {
                return Movement{-s, 0}
            } else {
                return Movement{s, 0}
            }

        } else {
            log.Fatal(err)
            log.Fatal(amount, "Couldn't parse as int")
        }
    } else {
        log.Fatal(command, "Didn't match expected regex")
    }
    return Movement{0, 0}
}

func parse_input(input_path string) []Movement{
    file, err := os.Open(input_path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var lines []Movement
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, parse_command(scanner.Text()))
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return lines

}

func main() {
    lines := parse_input("../inputs/input.in")
    fmt.Println(multiply_indexes(get_position(lines)))

}

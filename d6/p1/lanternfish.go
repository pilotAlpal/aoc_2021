package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "log"
    "strings"
)

func parse_input(input_path string) []int {
    file, err := os.Open(input_path)
    var ages []int
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scanner.Scan()
    str_values := strings.Split(scanner.Text(), ",")
    for _, v := range str_values {
        if s, err := strconv.Atoi(v); err == nil {
            ages = append(ages, s)
        } else {
            log.Fatal(err)
            log.Fatal("Could not convert to int")
        }
    }
    return ages
}

func get_next_population(current_pop []int) []int {
    var next_pop []int
    new_siblings := 0
    for _, v := range current_pop {
        if v == 0 {
            next_pop = append(next_pop, 6)
            new_siblings++
        } else {
            next_pop = append(next_pop, v - 1)
        }
    }
    for i := 0; i < new_siblings; i++ {
        next_pop = append(next_pop, 8)
	}
    return next_pop
}

func get_nth_day_population(days int, initial_pop []int) []int {
    current_pop := initial_pop
    for i := 0; i < days; i++ {
        current_pop = get_next_population(current_pop)
	}
    return current_pop

}

func main() {
    input_ages := parse_input("../inputs/input.in")
    nth_day_pop := get_nth_day_population(80, input_ages)
    fmt.Println(len(nth_day_pop))
}


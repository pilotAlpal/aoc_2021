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
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scanner.Scan()
    str_values := strings.Split(scanner.Text(), ",")
    timers := make([]int, 9)
    for _, v := range str_values {
        if s, err := strconv.Atoi(v); err == nil {
            timers[s]++
        } else {
            log.Fatal(err)
            log.Fatal("Could not convert to int")
        }
    }
    return timers
}

func get_next_population(current_pop *[]int) {
    zeros := (* current_pop)[0]
    for i, _ := range (* current_pop)[0:8] {
        (* current_pop)[i] = (* current_pop)[i + 1]
    }
    (* current_pop)[6] += zeros
    (* current_pop)[8] = zeros
}

func get_nth_day_population(days int, initial_pop *[]int) {
    for i := 0; i < days; i++ {
        get_next_population(initial_pop)
	}
}

func count_population(population []int) int {
    lanternfishes := 0
    for _, v := range population {
        lanternfishes += v
    }
    return lanternfishes
}
func main() {
    population := parse_input("../inputs/input.in")
    p := &population
    get_nth_day_population(256, p)
    fmt.Println(count_population(population))

}


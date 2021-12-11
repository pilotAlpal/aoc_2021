package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "log"
    "strings"
)

type CrabsMap struct {
	Crabs []int
	FuelConsumption []int
}

func get_min_position(inp []int) int {
    min := inp[0]
    p := 0
    for i, v := range inp {
        if v < min {
            min = v
            p = i
        }
    }
    return p
}


func (cm CrabsMap) get_consumption_to_i(initial_pos, incr int) int {
    sum := 0
    dist:= 1
    for i:= initial_pos + incr; 0 <= i && i <= len(cm.Crabs); i+= incr {
       sum += dist * cm.Crabs[i]
       dist++
    }
    return sum
}

func (cm CrabsMap) include_crab(c_pos int) CrabsMap {
    for i := len(cm.FuelConsumption) - 1; i <= c_pos; i++ {
        cm.Crabs = append(cm.Crabs, 0)
        cm.FuelConsumption = append(cm.FuelConsumption, cm.get_consumption_to_i(i, -1))
    }

    cm.Crabs[c_pos]++

    for i := 0; i < c_pos; i++ {
        cm.FuelConsumption[i] += c_pos - i
    }

    for i := c_pos + 1; i < len(cm.FuelConsumption); i++ {
        cm.FuelConsumption[i] += i - c_pos
    }
    return cm
}

func parse_input(input_path string) CrabsMap {
    cm := CrabsMap{}
    file, err := os.Open(input_path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scanner.Scan()
    for _, w := range strings.Split(scanner.Text(), ",") {
        if s, err := strconv.Atoi(w); err == nil {
            cm = cm.include_crab(s)
        } else {
            log.Fatal(err)
            log.Fatal("Could not convert to int")
        }
    }
    return cm
}

func main() {
    c_map := parse_input("../inputs/input.in")
    cheapest := get_min_position(c_map.FuelConsumption)
    fmt.Println(c_map.FuelConsumption[cheapest])
}


package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func main() {
    file, err := os.Open("../inputs/input.in")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var lines []int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        if s, err := strconv.Atoi(scanner.Text()); err == nil {
            lines = append(lines, s)
	    }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    windows := make([]int, len(lines)-2)
    for i, _ := range windows {
        windows[i] = lines[i] + lines[i+1] + lines[i+2]
	}
    last_depth := windows[0]
    depths := windows[1:]
    increases_count := 0
    for _, v := range depths {
        if v > last_depth {
		    increases_count++
	    }
        last_depth = v
	}
    fmt.Println(increases_count)

}

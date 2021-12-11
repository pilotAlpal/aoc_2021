package main

import (
    "fmt"
    "os"
    "bufio"
    "regexp"
    "log"
    "strings"
)

func line_matches_regexp(in_l string) bool {
    expected_re := `([a-g]{2,7} ){10} | ([a-g]{2,7} *){4}`
    re := regexp.MustCompile(expected_re)
    return re.Match([]byte(in_l))
}

func get_target_displays(displays []string) int {
    tgt_d := 0
    for _, v := range displays {
        lv := len(v)
        if lv == 2 || lv == 3 || lv == 4 || lv == 7 {
            tgt_d++
        }
    }
    return tgt_d
}
func parse_input(input_path string) ([][]string, [][]string, int) {
    file, err := os.Open(input_path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    var segments, displays [][]string
    t_displays := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        l := scanner.Text()
        if line_matches_regexp(l) {
            l_args := strings.Split(l, "|")
            lth_displays := strings.Split(l_args[1], " ")
            segments = append(segments, strings.Split(l_args[0], " "))
            displays = append(displays, lth_displays)
            t_displays += get_target_displays(lth_displays)
        } else {
            log.Fatal(l, "Didn't match expected regex")
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return segments, displays, t_displays


}

func main() {
    _, _, target_displays := parse_input("../inputs/input.in")
    fmt.Println(target_displays)
}

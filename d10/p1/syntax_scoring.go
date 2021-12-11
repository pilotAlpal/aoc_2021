package main

import (
    "fmt"
    "regexp"
    "os"
    "log"
    "bufio"
)

type Stack struct {
    Elems []rune
    NItems int
}

func (s Stack) push(e rune) (Stack) {
    return Stack{append(s.Elems, e), s.NItems + 1}
}

func (s Stack) is_empty() bool {
    return s.NItems == 0
}

func (s Stack) pop() (Stack, rune) {
    top_e := s.Elems[s.NItems - 1]
    return Stack{s.Elems[:s.NItems - 1], s.NItems - 1}, top_e
}

func line_matches_regexp(in_l string) bool {
    expected_re := `[\<\>\{\}\[\]\(\)]+`
    re := regexp.MustCompile(expected_re)
    return re.Match([]byte(in_l))
}

func opens_chunk(e rune) bool {
    if e == rune('<') || e == rune('{') || e == rune('(') || e == rune('[') {
        return true
    }
    return false
}

func get_chunk_open(e rune) rune {
    if e == rune('>') {
        return rune('<')
    }
    if e == rune('}') {
        return rune('{')
    }
    if e == rune(')') {
        return rune('(')
    }
    return rune('[')
}

func get_char_score(e rune) int {
    if e == rune('>') {
        return 25137
    }
    if e == rune('}') {
        return 1197
    }
    if e == rune(')') {
        return 3
    }
    return 57
}

func process_line(line string) (bool, rune) {
    s := Stack{}
    for _, c := range line {
        if opens_chunk(c) {
            s = s.push(c)
        } else {
            if s.is_empty() {
                return false, rune(c)
            } else {
                var t_e rune
                s, t_e = s.pop()
                if get_chunk_open(c) != t_e {
                    return false, rune(c)
                }
            }
        }
    }
    return true, rune(0)
}

func parse_input(input_path string) int {
    file, err := os.Open(input_path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)

    score := 0
    for scanner.Scan() {
        l := scanner.Text()
        if line_matches_regexp(l) {
            if incorrupted, br_c := process_line(l); !incorrupted {
                score += get_char_score(br_c)
            }
        } else {
            log.Fatal(l, "Didn't match expected regex")
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return score
}

func main() {
    score := parse_input("../inputs/input.in")
    fmt.Println(score)
}


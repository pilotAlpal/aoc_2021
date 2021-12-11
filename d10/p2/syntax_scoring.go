package main

import (
    "fmt"
    "regexp"
    "os"
    "log"
    "bufio"
    "sort"
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

func (s Stack) get_closure() []rune {
    var closure []rune
    for !s.is_empty() {
        var t_e rune
        s, t_e = s.pop()
        closure = append(closure, get_chunk_closure(t_e))
    }
    return closure
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

func get_chunk_closure(e rune) rune {
    if e == rune('<') {
        return rune('>')
    }
    if e == rune('{') {
        return rune('}')
    }
    if e == rune('(') {
        return rune(')')
    }
    return rune(']')
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
        return 4
    }
    if e == rune('}') {
        return 3
    }
    if e == rune(')') {
        return 1
    }
    return 2
}

func get_score(closure []rune) int {
    score := 0
    for _, c := range closure {
        score = (5 * score) + get_char_score(c)
    }
    return score
}


func process_line(line string) (bool, []rune) {
    s := Stack{}
    for _, c := range line {
        if opens_chunk(c) {
            s = s.push(c)
        } else {
            var bc []rune
            if s.is_empty() {
                return false, append(bc, rune(c))
            } else {
                var t_e rune
                s, t_e = s.pop()
                if get_chunk_open(c) != t_e {
                    return false, append(bc, rune(c))
                }
            }
        }
    }
    return true, s.get_closure()
}

func parse_input(input_path string) []int {
    file, err := os.Open(input_path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)

    var scores []int
    var closures [][]rune
    for scanner.Scan() {
        l := scanner.Text()
        if line_matches_regexp(l) {
            if incorrupted, br_c := process_line(l); incorrupted {
                closures = append(closures, br_c)
                scores = append(scores, get_score(br_c))
            }
        } else {
            log.Fatal(l, "Didn't match expected regex")
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    sort.Ints(scores)
    return scores
}

func main() {
    scores := parse_input("../inputs/input.in")
    fmt.Println(scores[len(scores) / 2])
}


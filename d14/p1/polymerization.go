package main

import (
    "fmt"
    "regexp"
    "strings"
    "os"
    "log"
    "bufio"
    "math"

)

const POL_RE = `[A-Z]+`
const RULES_RE = `[A-Z]{2} -> [A-Z]`

type CharsList struct {
    V rune
    Next *CharsList
}

func (ll *CharsList) init(f_v rune) *CharsList {
    ll_item := CharsList{f_v, nil}
    return &ll_item
}

func (ll *CharsList) push(n_v rune) {
    if ll != nil {
        ll_item := CharsList{n_v, nil}
        current := ll
        for current.Next != nil {
            current = current.Next
        }
        current.Next = &ll_item
    }
}

func (ll *CharsList) take_steps(rules map[string]rune, steps int) {
    for s := 0; s < steps; s++ {
        prev_char := ll
        current_char := ll.Next
        for current_char != nil {
            if token, ok := rules[string([]rune{prev_char.V, current_char.V})]; ok {
                n_item := CharsList{token, current_char}
                prev_char.Next = &n_item
            }
            prev_char = current_char
            current_char = current_char.Next
        }
    }
}

func line_matches_regexp(in_l, expected_re string) bool {
    re := regexp.MustCompile(expected_re)
    return re.Match([]byte(in_l))
}

func (ll *CharsList) get_char_ocurrences() map[rune]int {
    ocurrences := make(map[rune]int)
    current_char := ll
    for current_char != nil {
        if count, ok := ocurrences[current_char.V]; ok {
            ocurrences[current_char.V] = count + 1
        } else {
            ocurrences[current_char.V] = 1
        }
        current_char = current_char.Next
    }
    return ocurrences
}

func (ll *CharsList) get_most_and_least_repeated_ocurrences() (int, int) {
    ocurrences := ll.get_char_ocurrences()
    max_rep := 0
    min_rep := math.MaxInt64
    for _, v := range ocurrences {
        if v > max_rep {
            max_rep = v
        }
        if v < min_rep {
            min_rep = v
        }
    }
    return max_rep, min_rep
}

func polymer_2_list(polimer string) *CharsList {
    var cl *CharsList
    cl = cl.init(rune(polimer[0]))
    for _, c := range polimer[1:] {
        cl.push(rune(c))
    }
    return cl

}

func (ll *CharsList) print_list() {
    c_item := ll
    for c_item != nil {
        fmt.Print(string(c_item.V))
        c_item = c_item.Next
    }
    fmt.Println("")
}

func parse_input(input_path string) (*CharsList, map[string]rune) {
    var polymer *CharsList
    rules := make(map[string]rune)
    file, err := os.Open(input_path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scanner.Scan()
    l := scanner.Text()
    if line_matches_regexp(l, POL_RE) {
        polymer = polymer_2_list(l)
        scanner.Scan()
        for scanner.Scan()  {
            l = scanner.Text()
            if line_matches_regexp(l, RULES_RE) {
                rule_splitted := strings.Split(l, " -> ")
                rules[rule_splitted[0]] = rune(rule_splitted[1][0])
            } else {
                log.Fatal(l, "Didn't match expected regex", RULES_RE)
            }
        }
    } else {
        log.Fatal(l, "Didn't match expected regex", POL_RE)
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return polymer, rules
}

func main() {
    polymer, rules := parse_input("../inputs/input.in")
    polymer.take_steps(rules, 10)
    max, min := polymer.get_most_and_least_repeated_ocurrences()
    fmt.Println(max - min)
}


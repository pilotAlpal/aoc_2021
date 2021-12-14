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

func line_matches_regexp(in_l, expected_re string) bool {
    re := regexp.MustCompile(expected_re)
    return re.Match([]byte(in_l))
}

func increase_segment_count(polymer map[[2]rune]int, segment [2]rune, incr int) map[[2]rune]int{
        if _, ok := polymer[segment]; ok {
            polymer[segment]+= incr
        } else {
            polymer[segment] = incr
        }
        return polymer
}

func decrease_segment_count(polymer map[[2]rune]int, segment [2]rune, decrease int) map[[2]rune]int {
        if _, ok := polymer[segment]; ok {
            polymer[segment]-= decrease
            if polymer[segment] <= 0 {
                delete(polymer, segment)
            }
        }
        return polymer
}

func add_segments(input string, polymer map[[2]rune]int) map[[2]rune]int {
    prev := rune(input[0])
    for _, c := range input[1:] {
        seg_k := [2]rune{prev, c}
        polymer = increase_segment_count(polymer, seg_k, 1)
        prev = rune(c)
    }
    return polymer
}

func take_steps(polymer map[[2]rune]int, rules map[[2]rune]rune, steps int) map[[2]rune]int {

    for s := 0; s < steps; s++ {
        sth_pol_keys := make([][2]rune, len(polymer))
        sth_pol_values := make([]int, len(polymer))
        for k, v := range polymer {
            sth_pol_keys = append(sth_pol_keys, k)
            sth_pol_values = append(sth_pol_values, v)
        }
        for i, k := range sth_pol_keys {
            if r, ok := rules[k]; ok {
                polymer = decrease_segment_count(polymer, k, sth_pol_values[i])
                polymer = increase_segment_count(polymer, [2]rune{k[0], r}, sth_pol_values[i])
                polymer = increase_segment_count(polymer, [2]rune{r, k[1]}, sth_pol_values[i])
            }
        }
    }
    return polymer
}

func parse_input(input_path string) (map[[2]rune]int, map[[2]rune]rune, rune, rune) {
    polymer := make(map[[2]rune]int)
    rules := make(map[[2]rune]rune)
    file, err := os.Open(input_path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scanner.Scan()
    l := scanner.Text()
    var first, last rune
    if line_matches_regexp(l, POL_RE) {
        polymer = add_segments(l, polymer)
        first, last = rune(l[0]), rune(l[len(l) - 1])
        scanner.Scan()
        for scanner.Scan()  {
            l = scanner.Text()
            if line_matches_regexp(l, RULES_RE) {
                rule_splitted := strings.Split(l, " -> ")
                ori0, ori1, dest := rune(rule_splitted[0][0]), rune(rule_splitted[0][1]), rune(rule_splitted[1][0])
                rules[[2]rune{ori0, ori1}] = dest
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
    return polymer, rules, first, last
}

func print_segments(segments map[[2]rune]int) {
    for k, v := range segments {
        fmt.Println(string(k[0]), string(k[1]), ":", v)
    }
}

func print_rules(rules map[[2]rune]rune) {
    for k, v := range rules {
        fmt.Println(string(k[0]), string(k[1]), "->", string(v))
    }
}

func print_ocurrences(ocurrences map[rune]int) {
    for k, v := range ocurrences {
        fmt.Println(string(k), "=", v)
    }
}

func get_ocurrences(segments map[[2]rune]int, f, l rune) map[rune]int {
    ocurrences := make(map[rune]int)
    ocurrences[f] = 1
    ocurrences[l] = 1
    for k, v := range segments {
        for _, p := range k {
            if _, ok := ocurrences[p]; ok {
                ocurrences[p] += v
            } else {
                ocurrences[p] = v
            }
        }

    }
    for k, v := range ocurrences {
        ocurrences[k] = v / 2
    }
    return ocurrences
}

func get_most_and_least_repeated(ocurrences map[rune]int) (int, int) {
    max_rep := 0
    min_rep := math.MaxInt64
//    print_ocurrences(ocurrences)
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


func main() {
    polymer, rules, f, l := parse_input("../inputs/input.in")
//    print_rules(rules)
    polymer = take_steps(polymer, rules, 40)
//    print_segments(polymer)
    max, min := get_most_and_least_repeated(get_ocurrences(polymer, f, l))
    fmt.Println(max - min)
}


package main

import (
    "fmt"
    "os"
    "bufio"
    "regexp"
    "log"
    "strings"
    "sort"
)

func line_matches_regexp(in_l string) bool {
    expected_re := `([a-g]{2,7} ){10} | ([a-g]{2,7} *){4}`
    re := regexp.MustCompile(expected_re)
    return re.Match([]byte(in_l))
}

func sort_string(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

func get_sequence(in string, segments_mapping map[rune]rune) string {
    s := ""
    for _, v := range in {
        s = s + string(segments_mapping[rune(v)])
    }
    return s
}

func get_mapping(segments []string) map[string]int {
    digits_mapping := make(map[string]int)
    segments_counter := make(map[rune]int)
    sequence_of_4 := ""
    for _, v := range segments {
        lv := len(v)
        sorted_v := sort_string(v)
        if lv == 2 {
            digits_mapping[sorted_v] = 1
        } else if lv == 3 {
            digits_mapping[sorted_v] = 7
        } else if lv == 4 {
            digits_mapping[sorted_v] = 4
            sequence_of_4 = sorted_v
        } else if lv == 7 {
            digits_mapping[sorted_v] = 8
        }
        for _, char := range sorted_v {
            _, ok := segments_counter[char]
            if !ok {
                segments_counter[char] = 1
            } else {
                segments_counter[char]++
            }
        }
    }
    segments_map := make(map[rune]rune)
    for k, v := range segments_counter {
        if v == 4 {
            segments_map[rune('e')] = rune(k)
        } else if v == 6 {
            segments_map[rune('b')] = rune(k)
        } else if v == 7 {
            if strings.ContainsRune(sequence_of_4, rune(k)) {
                segments_map[rune('d')] = rune(k)
            } else {
                segments_map[rune('g')] = rune(k)
            }
        } else if v == 8 {
            if strings.ContainsRune(sequence_of_4, rune(k)) {
                segments_map[rune('c')] = rune(k)
            } else {
                segments_map[rune('a')] = rune(k)
            }
        } else if v == 9 {
            segments_map[rune('f')] = rune(k)
        } else {
            log.Fatal(v, k, "Unexpected length")
        }
    }

    digits_mapping[sort_string(get_sequence(`abcefg`, segments_map))] = 0
    digits_mapping[sort_string(get_sequence(`acdeg`, segments_map))] = 2
    digits_mapping[sort_string(get_sequence(`acdfg`, segments_map))] = 3
    digits_mapping[sort_string(get_sequence(`abdfg`, segments_map))] = 5
    digits_mapping[sort_string(get_sequence(`abdefg`, segments_map))] = 6
    digits_mapping[sort_string(get_sequence(`abcdfg`, segments_map))] = 9
    return digits_mapping

}

func get_digits(displays []string, mapping map[string]int) []int {
   digits := make([]int, 4)
   for i, v := range displays {
        digits[i] = mapping[sort_string(v)]
   }
   return digits
}

func sum(digits []int) int {
    sum := 0
    shift := 1000
    for _, v := range digits {
        sum += v * shift
        shift = shift / 10
    }
    return sum
}
func parse_input(input_path string) ([][]int, int) {
    file, err := os.Open(input_path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    var outputs [][]int
    outputs_sum := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        l := scanner.Text()
        if line_matches_regexp(l) {
            l_args := strings.Split(l, " | ")
            lth_segments := strings.Split(l_args[0], " ")
            mapping := get_mapping(lth_segments)
            lth_displays := strings.Split(l_args[1], " ")
            digits := get_digits(lth_displays, mapping)
            outputs_sum += sum(digits)
            outputs = append(outputs, digits)
        } else {
            log.Fatal(l, "Didn't match expected regex")
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return outputs, outputs_sum


}

func main() {
    _, target_displays := parse_input("../inputs/input.in")
    fmt.Println(target_displays)
}

package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
	"regexp"
    "errors"
//    "strings"
    "strconv"
)

type BitsCounter struct {
	Zeros int
	Ones int
}

func get_bits_counter(b rune) (BitsCounter, error) {
    if b == '0' {
        return BitsCounter{1, 0}, nil
    } else if b == '1' {
        return BitsCounter{0, 1}, nil
    } else {
        log.Fatal(b, "Couldn't parse as int")
        return BitsCounter{}, errors.New("Couldn't parse as int")
    }
}

func add_bits_counters(bc0, bc1 BitsCounter) BitsCounter {
    return BitsCounter{bc0.Zeros + bc1.Zeros, bc0.Ones + bc1.Ones}
}

func add_bits_to_count(bits_c []BitsCounter, input string) []BitsCounter {
    for i, v := range bits_c {
        ith_bc, err := get_bits_counter(rune(input[i]))
        if err != nil {
            log.Fatal(err)
        }
        bits_c[i] = add_bits_counters(v, ith_bc)
    }
    for i:= len(bits_c); i < len(input); i++ {
        ith_bc, err := get_bits_counter(rune(input[i]))
        if err != nil {
            log.Fatal(err)
        }
        bits_c = append(bits_c, ith_bc)
    }
    return bits_c
}

func matches_regexp(input string) bool {
	re := regexp.MustCompile(`[0|1]+`)
    return re.Match([]byte(input))
}

func get_gamma_and_epsilon(b_counter []BitsCounter) (string, string){
    g := ""
    e := ""
    for _, v := range b_counter {
        if v.Zeros > v.Ones {
            g = g + "0"
            e = e + "1"
        } else {
            g = g + "1"
            e = e + "0"
        }
   }
   return g, e
}

func parse_input(input_path string) []string {
    file, err := os.Open(input_path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var d_report []string
    for ;scanner.Scan(); matches_regexp(scanner.Text()) {
        d_report = append(d_report, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return d_report

}

func get_most_repeated_ith_item(bin_strs []string, offset int) string {
    zeros := 0
    ones := 0
    for _, v := range bin_strs {
        if string(v[offset]) == "0" {
            zeros++
        } else {
            ones++
        }
    }
    if ones >= zeros {
        return "1"
    }
    return "0"
}

func get_oxygen_rating(most_commons []string) string{
    offset := 0
    for len(most_commons) > 1 {
        most_repeated := get_most_repeated_ith_item(most_commons, offset)
        var mc []string
        for _, w := range most_commons {
            if string(w[offset]) == most_repeated {
                mc = append(mc, w)
            }
        }
        most_commons = mc
        offset++
    }
    return most_commons[0]
}

func get_scrubber_rating(least_common []string) string{
    offset := 0
    for len(least_common) != 1 {
        most_repeated := get_most_repeated_ith_item(least_common, offset)
        var lc []string
        for _, w := range least_common {
            if string(w[offset]) != most_repeated {
                lc = append(lc, w)
            }
        }
        least_common = lc
        offset++

    }
    return least_common[0]
}

func binstr_to_int(binstr string) (int, error) {
   if s, err := strconv.ParseInt(binstr, 2, 64); err == nil {
            return int(s), nil

   } else {
        log.Fatal(err)
        log.Fatal(binstr, "Couldn't parse as int")
        return -1, errors.New("Couldn't parse as int")
   }
}


func get_life_support_rating(ox_r, sc_r string) int {
   var ox_v, sc_v int
   if s, err := binstr_to_int(ox_r); err == nil {
        ox_v = int(s)

   } else {
        log.Fatal(err)
        log.Fatal(ox_r, "Couldn't parse as int")
        return -1
   }
   if s, err := binstr_to_int(sc_r); err == nil {
        sc_v = int(s)

   } else {
        log.Fatal(err)
        log.Fatal(sc_r, "Couldn't parse as int")
        return -1
   }
   return ox_v * sc_v

}

func main() {
    diagnostics := parse_input("../inputs/input.in")
    ox_r := get_oxygen_rating(diagnostics)
    sc_r := get_scrubber_rating(diagnostics)
    fmt.Println(get_life_support_rating(ox_r, sc_r))
}

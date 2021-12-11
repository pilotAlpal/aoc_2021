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

func get_gamma_and_epsilon(b_counter []BitsCounter) (int, int){
    g := ""
    e := ""
    var gamma, epsilon int
    for _, v := range b_counter {
        if v.Zeros > v.Ones {
            g = g + "0"
            e = e + "1"
        } else {
            g = g + "1"
            e = e + "0"
        }
   }
   if s, err := strconv.ParseInt(e, 2, 64); err == nil {
            epsilon = int(s)

   } else {
        log.Fatal(err)
        log.Fatal(e, "Couldn't parse as int")
   }
   if s, err := strconv.ParseInt(g, 2, 64); err == nil {
            gamma = int(s)

   } else {
        log.Fatal(err)
        log.Fatal(g, "Couldn't parse as int")
   }
   return gamma, epsilon
}
func parse_input(input_path string) []BitsCounter {
    file, err := os.Open(input_path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var bc []BitsCounter
    for ;scanner.Scan(); matches_regexp(scanner.Text()) {
        bc = add_bits_to_count(bc, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return bc

}

func main() {
    bcount := parse_input("../inputs/input.in")
    g, e := get_gamma_and_epsilon(bcount)
    fmt.Println(g * e)
}

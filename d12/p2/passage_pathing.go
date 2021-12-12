package main

import (
    "fmt"
    "regexp"
    "os"
    "log"
    "strings"
    "bufio"
)

func line_matches_regexp(in_l string) bool {
    expected_re := `[[a-z]+|[A-Z]+]-[[a-z]+|[A-Z]+]`
    re := regexp.MustCompile(expected_re)
    return re.Match([]byte(in_l))
}

func add_path(g map[string][]string, x, y string) map[string][]string {
    g[x] = append(g[x], y)
    g[y] = append(g[y], x)
    return g
}

func get_points(line string) (string, string) {
    points := strings.Split(line, "-")
    return points[0], points[1]
}

func parse_input(input_path string) map[string][]string {
    file, err := os.Open(input_path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)

    var graph = make(map[string][]string)
    for scanner.Scan() {
        l := scanner.Text()
        if line_matches_regexp(l) {
            p0, p1 := get_points(l)
            graph = add_path(graph, p0, p1)
        } else {
            log.Fatal(l, "Didn't match expected regex")
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return graph
}

func is_small_cave(cave string) bool {
    expected_re := `[a-z]+`
    re := regexp.MustCompile(expected_re)
    return re.Match([]byte(cave))
}

func include_paths(current_paths, new_paths [][]string) [][]string {
    if len(current_paths) == 0 {
        return new_paths
    }
    if len(new_paths) == 0 {
        return current_paths
    }
    current_len := len(current_paths)
    joined := make([][]string, current_len + len(new_paths))
    _ = copy(joined, current_paths)
    _ = copy(joined[current_len:], new_paths)
    return joined
}

func find_paths(graph map[string][]string, start, end string, current_path []string, smalls_visited map[string]bool, paths int, small_repeated bool) int {
    current_path = append(current_path, start)
    if start == end {
        return paths + 1
    }
    for _, p := range graph[start] {
        if _, visited := smalls_visited[p]; !visited {
            if is_small_cave(p) {
                smalls_visited[p] = true
            }
            paths = find_paths(graph, p, end, current_path, smalls_visited, paths, small_repeated)
            delete(smalls_visited, p)
        } else if !small_repeated && p != end && p!= "start" {
            paths = find_paths(graph, p, end, current_path, smalls_visited, paths, true)
        }
    }
    return paths
}

func get_paths(graph map[string][]string, start, end string) int {
    scv := make(map[string]bool)
    scv[start] = true
    return find_paths(graph, start, end, []string{}, scv, 0, false)
}

func main() {
    graph := parse_input("../inputs/input.in")
    fmt.Println(graph)
    paths := get_paths(graph, "start", "end")
    fmt.Println(paths)
}


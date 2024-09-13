package part2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadInput2(fileName string) (map[string]map[string]int, []string, error) {
    file, err := os.Open(fileName)
    if err != nil {
        return nil, nil, err
    }
    defer file.Close()

    distances := make(map[string]map[string]int)
    locationsMap := make(map[string]bool)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, " ")
        from := parts[0]
        to := parts[2]
        dist, _ := strconv.Atoi(parts[4])

        if distances[from] == nil {
            distances[from] = make(map[string]int)
        }
        if distances[to] == nil {
            distances[to] = make(map[string]int)
        }

        distances[from][to] = dist
        distances[to][from] = dist

        locationsMap[from] = true
        locationsMap[to] = true
    }

    if err := scanner.Err(); err != nil {
        return nil, nil, err
    }

    var locations []string
    for loc := range locationsMap {
        locations = append(locations, loc)
    }

    return distances, locations, nil
}

func permutations(locations []string) [][]string {
    var results [][]string
    permute(locations, []string{}, &results)
    return results
}

func permute(locations []string, current []string, results *[][]string) {
    if len(locations) == 0 {
        *results = append(*results, append([]string(nil), current...))
        return
    }
    for i, loc := range locations {
        next := append([]string{}, locations[:i]...)
        next = append(next, locations[i+1:]...)
        permute(next, append(current, loc), results)
    }
}

func LongestRoute(locations []string, dist map[string]map[string]int) int {
    perms := permutations(locations)
    maxDist := 0

    for _, route := range perms {
        totalDist := 0
        for i := 0; i < len(route)-1; i++ {
            totalDist += dist[route[i]][route[i+1]]
        }
        if totalDist > maxDist {
            maxDist = totalDist
        }
    }
    return maxDist
}

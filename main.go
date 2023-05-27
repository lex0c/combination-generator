package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "sync"
    "math/rand"
    "time"
)

func combinations(arr []string, buffer []string, i, n, bufIndex int, out chan []string, wg *sync.WaitGroup) {
    defer wg.Done()

    if bufIndex == len(buffer) {
        combination := make([]string, len(buffer))
        copy(combination, buffer)
        out <- combination
    } else {
        for j := i; j < n; j++ {
            buffer[bufIndex] = arr[j]
            wg.Add(1)
            combinations(arr, buffer, j+1, n, bufIndex+1, out, wg)
        }
    }
}

func generateAllCombinations(arr []string, r int) <-chan []string {
    out := make(chan []string)

    wg := &sync.WaitGroup{}

    go func() {
        defer close(out)

        buffer := make([]string, r)

        wg.Add(1)

        go combinations(arr, buffer, 0, len(arr), 0, out, wg)

        wg.Wait()
    }()

    return out
}

func shuffleArray(arr []string) {
    rand.Seed(time.Now().UnixNano())

    rand.Shuffle(len(arr), func(i, j int) {
        arr[i], arr[j] = arr[j], arr[i] // Swap elements at indices i and j
    })
}

func main() {
    if len(os.Args) != 4 {
        fmt.Println("Provide an input file, output file, and combination length as command-line arguments")
        os.Exit(1)
    }

    combinationLength, err := strconv.Atoi(os.Args[3])

    if err != nil {
        fmt.Printf("Error: combination length must be a valid integer, not %s: %v\n", os.Args[3], err)
        os.Exit(1)
    }

    inFile, err := os.Open(os.Args[1])

    if err != nil {
        fmt.Printf("Error opening input file %s: %v\n", os.Args[1], err)
        os.Exit(1)
    }

    defer inFile.Close()

    scanner := bufio.NewScanner(inFile)

    var linesRaw []string

    for scanner.Scan() {
        linesRaw = append(linesRaw, scanner.Text())
    }

    lines := make([]string, len(linesRaw))
    copy(lines, linesRaw)
    shuffleArray(lines)

    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading input file %s: %v\n", os.Args[1], err)
        os.Exit(1)
    }

    if _, err := os.Stat(os.Args[2]); !os.IsNotExist(err) {
        fmt.Printf("Output file %s already exists.\n", os.Args[2])
        os.Exit(1)
    }

    outFile, err := os.Create(os.Args[2])

    if err != nil {
        fmt.Printf("Error creating output file %s: %v\n", os.Args[2], err)
        os.Exit(1)
    }

    defer outFile.Close()

    for combination := range generateAllCombinations(lines, combinationLength) {
        strCombination := strings.Join(combination, " ")

        if _, err := outFile.WriteString(strCombination + "\n"); err != nil {
            fmt.Printf("Error writing to output file %s: %v\n", os.Args[2], err)
            os.Exit(1)
        }
    }

    fmt.Printf("Combinations saved to %s\n", os.Args[2])
}


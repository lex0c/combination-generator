package main

import (
    "testing"
    "sync"
    "strings"
)

func TestCombinations(t *testing.T) {
    arr := []string{"a", "b", "c"}

    buffer := make([]string, 2)

    out := make(chan []string)

    wg := &sync.WaitGroup{}

    go func() {
        defer close(out)
        wg.Add(1)
        go combinations(arr, buffer, 0, len(arr), 0, out, wg)
        wg.Wait()
    }()

    result := []string{}

    for combination := range out {
        result = append(result, combination...)
    }

    expected := "abacbc"

    if strings.Join(result, "") != expected {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}


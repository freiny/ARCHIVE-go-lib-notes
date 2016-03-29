package main

import (
    "os"
    "strings"
    "fmt"
    "bufio"
)

func main() {
    removeNewLines("main.go")
}

func removeNewLines(path string) {
    infile, err := os.Open(path)

    if err != nil {
        fmt.Println("Error opening file:", err)
    }

    scanner := bufio.NewScanner(infile)
    scanner.Split(bufio.ScanLines)

    for scanner.Scan() {
        line := scanner.Text()
        if strings.TrimSpace(string(line)) == "" {
        } else {
            // fmt.Println(strings.TrimSpace(string(line)))
            fmt.Println(line)
        }
    }

}

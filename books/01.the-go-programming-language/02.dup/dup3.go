package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    counts := make(map[string]int)
    
    for _, filename := range os.Args[1:] {
        data, err := ioutil.ReadFile(filename)
        
        if err != nil {
            _, err := fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
            if err != nil {
                continue 
            }
            continue
        }
        
        for _, line := range strings.Split(string(data), "\n") {
            counts[line]++
        }
        
        for line, n := range counts {
            fmt.Printf("%d\t%s", n, line)
        }
    }
}

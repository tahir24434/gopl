package main
import (
    "fmt"
    "./links"
    "log"    
    "os"
)

func breadthFirst(f func(item string) []string, worklist []string) {
    seen := make(map[string]bool)
    level := worklist
    for len(level) > 0 {
        var next_level []string
        for _, item := range level {
            if !seen[item] {
                seen[item] = true
                next_level = append(next_level, f(item)...)
            }
        }
        level = next_level
    }
}

func crawl(url string) []string {
    fmt.Println(url)
    list, err := links.Extract(url)
    if err != nil {
        log.Print(err)
    }
    return list
}

func main() {
    breadthFirst(crawl, os.Args[1:])
}

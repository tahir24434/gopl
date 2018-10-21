package main

import (
	"fmt"
)

func nonempty(strings []string) []string {
	e := 0
	for _, s := range strings {
		if s != "" {
			strings[e] = s
			e++
		}
	}
	return strings[:e]
}

func main() {
	strings := []string{"first", "", "second"}
	fmt.Printf("%q\n", nonempty(strings))
	fmt.Printf("%q\n", strings)
}

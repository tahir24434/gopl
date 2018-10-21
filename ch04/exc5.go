package main

import(
	"fmt"
)

func rem_adj(strings []string) []string {
	cursor := 0
	for _, s := range strings {
		if s != strings[cursor] {
			cursor++
		} 
		strings[cursor] = s
	}
	return strings[:cursor+1]
}

func rem_adj(strings []string) []string {
	out := strings[:0]
	i := 0
	for _, s := range strings {
		
	}
}

func main() {
	strings := []string{"1", "1", "2", "2", "4", "5", "5"}
	fmt.Printf("%q\n", rem_adj(strings))
	fmt.Printf("%q\n", strings)
}

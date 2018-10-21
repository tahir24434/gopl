package main
import (
	"fmt"
)

type LineCounter int

func (lc *LineCounter) Write(p []byte) (n int, err error) {
	count := 1
	for _, c := range p {
		if c == '\n' {
			count++
		}
	}
	*lc += LineCounter(count)
	return count, nil
}

func main() {
	var lc LineCounter
	b := "Mukhree"
	fmt.Fprintf(&lc, "Hello, %s\n bestest \n one \n", b)
	fmt.Println(lc)
}


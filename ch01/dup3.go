/*
dup above operate in a “streaming” modein which input is read and broken into lines as needed,so in principle these programs can handle an arbitrary amount of input.An alternative approach is to read the entire input into memory in one big gulp,split it into lines all at once, then process the lines.The following version, dup3, operates in that fashion.It introduces the function ReadFile (from the io/ioutilpackage), which readsthe entire contents of a named file, and strings.Split,which splits a string into a slice of substrings.(Split is the opposite of strings.Join, which wesaw earlier.)
*/

package main
import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	for _, file := range files {
		// func ReadFile(filename string) ([]byte, error)
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}


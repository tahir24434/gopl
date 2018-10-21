package main

import (
	"fmt"
	"bytes"	
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents an empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint64(x%64)
	return word < len(s.words) && s.words[word] & (1 << bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint64(x%64)
	for x > len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}"
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		} 
		for j := 0; j < 64; j++ {
			if word & (1 << uint(j)) != 0 {
				fmt.Fprintf(&buf, " %d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Return the number of elements
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		for j :=0; j < 64; j++ {
			if word & (1<<uint(j)) !=0 {
				count++
			}
		}
	}
	return count
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint64(x%64)
	if s.words[word] == 0 {
		return 
	}
	s.words[word] = s.words[word] & ^(1 << bit)
}

// Remove all the elements from the set
func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

// Return a copy of the set
func (s *IntSet) Copy() *IntSet {
	new := &IntSet{}
	for _, word := range s.words {
		new.words = append(new.words, word)
	}
	return new
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(19)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())
	fmt.Println(x.Has(9))
	fmt.Println(x.Len())
	x.Remove(9)
	fmt.Println(x.Has(9))

	new := y.Copy()
	fmt.Println(new)
}


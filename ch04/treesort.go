package main
import (
	"fmt"
)

type node struct {
	value int
	left, right *node
}

func add(t *node, val int) *node {
	if t == nil {
		// Equilant to return &node{value: val}
		t = new(node)
		t.value = val
	} else if val < t.value {
		t.left = add(t.left, val)
	} else {
		t.right = add(t.right, val)
	}
	return t
}

func inorder(t *node, values []int) []int {
	if t != nil {
		values = inorder(t.left, values)
		values = append(values, t.value)
		values = inorder(t.right, values)
	}
	return values 
}

func Sort(values []int) {
	var root *node
	for _, val := range values {
		root = add(root, val)
	}
	inorder(root, values[:0])
}

func main() {
	v := []int{5, 9, 11, 22, 1, 6, 9, 10}
	Sort(v)
	fmt.Printf("%v\n", v)
}

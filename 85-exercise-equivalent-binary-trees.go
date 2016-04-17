package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)

	var walk func(t *tree.Tree, ch chan int)
	walk = func(t *tree.Tree, ch chan int) {
		if t.Left != nil {
			walk(t.Left, ch)
		}
		ch <- t.Value
		if t.Right != nil {
			walk(t.Right, ch)
		}
	}
	walk(t, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v1 := range ch1 {
		if v2, ok := <-ch2; !(ok && v1 == v2) {
			return false
		}
	}
	if _, ok := <-ch2; ok {
		return false
	}

	return true
}

func main() {
	// Test of Walk
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}

	// Test of Same
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

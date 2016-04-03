package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"time"
)

func Walk(t *tree.Tree, ch chan int) {
	if t.Left == nil && t.Right == nil {
		ch <- t.Value
		return
	}
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	return
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		v1 := <-ch1
		v2 := <-ch2
		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	// Test of Walk
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	time.Sleep(100)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	// Test of Same
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

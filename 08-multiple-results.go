package main

import "fmt"

func swqp(x, y string) (string, string) {
	return x + y
}

func main() {
	a, b  := swqp("hello", "world")
	fmt.Println(a, b)
}

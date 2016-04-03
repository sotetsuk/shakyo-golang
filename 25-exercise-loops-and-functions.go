package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z, epsilon := x, 1e-10
	for i := 0; i < 100; i++ {
		z = z - (z*z-x)/(2*z)
		if d := math.Abs(z - x); d < epsilon {
			fmt.Printf("at %v\n", i)
			return z
		}
	}
	fmt.Println("@ over 99")
	return z
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Printf("Sqrt(%v):\n", i)
		fmt.Println(Sqrt(float64(i)))
	}
}

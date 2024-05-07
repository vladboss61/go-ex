package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Swap[T interface {}](a, b T) (T, T) {
    return b, a
}

// any == interface {}
func SwapAny[T any](a, b T) (T, T) {
    return b, a
}

func StartAbs() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

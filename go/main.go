package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	ds := [][]float64{
		{0, 0},
		{1, 2},
		{2, 4},
		{3, 6},
		{4, 8},
	}

	rate := 0.1
	weight := rand.Float64() * 10
	loss := 0.0

	for epoch := 0; epoch < 10; epoch++ {
		for input := range ds {
			x := ds[input][0]
			y := ds[input][1]
			yh := x * weight
			loss = math.Pow(y-yh, 2)
			gradient := 2 * x * (yh - y)
			weight -= rate * gradient
		}
		fmt.Printf("Epoch: %d Weight: %f Loss: %f\n", epoch, weight, loss)
	}
}

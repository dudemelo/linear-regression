package main

import (
	"fmt"
	"math"
)

func main() {
	train := [][2]float64{
		{0, 0},
		{1, 2},
		{2, 4},
		{3, 6},
		{4, 8},
	}
	learningRate := 0.1
	weight := 10.0
	for _, input := range train {
		fmt.Printf("=============== input[0]: %v\n", input[0])
		previousLoss := 0.0
		for i := 0; i < 10; i++ {
			output := input[0] * weight
			loss := math.Pow(input[1]-output, 2)
			if math.Abs(loss-previousLoss) < 0.00001 {
				break
			}
			previousLoss = loss
			fmt.Printf("output: %f expected: %f weight: %f loss: %f\n", output, input[1], weight, loss)
			gradient := -2 * input[0] * (input[1] - output)
			weight -= learningRate * gradient
		}
	}
}

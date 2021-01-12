package outliers

import (
	"math"
	"github.com/pavelerokhin/gostatistics/basics"
)

func ChauvenetFilter(data []float64) []float64 {
	output := data
	var currentNumerosity int

	for {
		currentNumerosity = len(output)
		output = chauvenetStep(output)

		if currentNumerosity == len(output) {
			break
		}
	}

	return output
}

func chauvenetStep(data []float64) []float64 {
	mean := basics.Mean(data)
	standardDeviation := basics.StandardDeviation(data)
	n := float64(len(data))

	var output []float64

	for _, value := range data {
		if n * math.Erfc((value - mean)/standardDeviation) > 0.5 {
			output = append(output, value)
		}
	}


	return output
}

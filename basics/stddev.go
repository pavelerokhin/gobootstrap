package basics

import "math"

func StandardDeviation(data []float64) float64 {
	mean := Mean(data)
	n := float64(len(data))

	var sum float64
	for _, value := range data {
		sum += math.Pow(value - mean, 2)
	}

	return math.Sqrt(sum/n)
}

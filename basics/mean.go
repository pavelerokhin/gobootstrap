package basics

func Mean(data []float64) float64 {
	var mean float64
	for _, value := range data {
		mean += value
	}
	return mean/float64(len(data))
}

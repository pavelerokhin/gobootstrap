package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/stat"
	"os"
	"sort"
)


func boot(data []float64, nBoot int, alpha float64) (float64, float64) {

	dataMean, err := stats.Mean(data)

	if err != nil {
		fmt.Printf("Error while calculating data mean")
		os.Exit(1)
	}

	bootSamples := bootSampling(data, nBoot)
	sort.Float64s(bootSamples)
	var bootDeltas []float64
	for _, s := range bootSamples {
		bootDeltas = append(bootDeltas, s - dataMean)
	}

	q1 := stat.Quantile((1-alpha)/2, stat.LinInterp, bootDeltas, nil)
	q2 := stat.Quantile(alpha+((1-alpha)/2), stat.LinInterp, bootDeltas, nil)

	return dataMean - q2, dataMean - q1
}

func bootSampling(data []float64, nBoot int) []float64 {

	var samples []float64
	sampleSize := len(data)

	for i := 0; i < nBoot; i++ {
		s, err := stats.Sample(data, sampleSize, true)
		if err!= nil {
			fmt.Printf("Error while sampling boot samples")
			os.Exit(1)
		}
		sMean := stat.Mean(s, nil)

		samples = append(samples,sMean)
	}
	return samples
}

func main() {
	data := []float64{163.84, 163.84, 327.68, 655.36, 1966.08, 2129.92, 2129.92, 1966.08, 1966.08, 1802.24, 1638.40, 1474.56, 1802.24, 1802.24,
		1966.08, 1474.56, 1638.40, 1802.24, 1474.56, 1638.40, 1638.40, 1638.40, 1638.40, 1638.40, 1638.40, 1638.40, 1474.56, 1638.40,
		1802.24, 1966.08, 1310.72, 1638.40, 1802.24, 1638.40, 1310.72, 1638.40, 1802.24, 1474.56, 1802.24, 1802.24, 1638.40, 1310.72,
		1802.24, 1638.40, 1310.72, 1474.56, 1638.40, 1474.56, 1146.88, 1474.56, 1310.72, 1310.72, 1474.56, 1310.72, 1638.40, 1638.40,
		1638.40, 1802.24, 1802.24, 1638.40, 1638.40, 1802.24, 1638.40, 1310.72, 1638.40, 1474.56, 1474.56, 1966.08, 1310.72, 1966.08,
		1638.40, 1802.24, 1802.24, 1638.40, 1638.40, 1802.24, 1802.24, 1638.40, 1638.40, 1802.24, 1638.40, 1802.24, 1474.56, 1638.40,
		1474.56, 983.04, 1638.40, 1802.24, 1802.24, 1638.40, 1802.24, 1802.24, 1802.24, 1966.08, 1966.08, 1638.40, 2129.92, 1638.40,
		1802.24, 1802.24, 1474.56}

	l, h := boot(data, 10000, 0.95)
	fmt.Printf("%v, %v", l, h)
}

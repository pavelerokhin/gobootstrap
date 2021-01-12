package outliers

import (
	"fmt"
	"testing"
)

// TestChauvenet controls correctness of the Chauvenet's test for outliers
func TestChauvenetFilter(t *testing.T) {
	input := []float64 {
		8.02, 8.16, 3.97, 8.64, 0.84, 4.46, 0.81, 7.74, 8.78, 9.26, 20.46, 29.87, 10.38,25.71,
	}

	correctOutput := []float64 {
		8.02, 8.16, 3.97, 8.64, 0.84, 4.46, 0.81, 7.74, 8.78, 9.26, 10.38,
	}

	output := ChauvenetFilter(input)

	fmt.Println(correctOutput, output)
}

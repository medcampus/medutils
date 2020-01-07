package medutils

import (
	"fmt"
	"math"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	fmt.Println(float64(round(num*output)) / output)
	return float64(round(num*output)) / output
}

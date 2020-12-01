package lib

import (
	"fmt"
	"math"
)

func ManhattanDistance(a []float64, b []float64) float64 {
	var n int
	var s float64
	if len(a) != len(b) {
		fmt.Println("Size differ")
		return 0
	}
	n = len(a)
	s = 0
	for i := 0; i < n; i += 1 {
		s += math.Abs(b[i] - a[i])
	}
	return float64(s)
}

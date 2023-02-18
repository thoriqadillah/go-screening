package array

import (
	"math"
)

func Max(numbers []int) int {
	max := math.MinInt
	for _, v := range numbers {
		if v > max {
			max = v
		}
	}

	return max
}

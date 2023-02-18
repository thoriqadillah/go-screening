package array

import (
	"math"
)

func GetMaxNumber(numbers []int) int {
	max := math.MinInt
	for _, v := range numbers {
		if v > max {
			max = v
		}
	}

	return max
}

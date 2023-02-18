package array

import "testing"

func TestGetMaxNumber(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	max := GetMaxNumber(arr)

	if max != 5 {
		t.Errorf("Expected result: %d. Actual result: %d", 5, max)
	}
}

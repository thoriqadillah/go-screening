package array

import "testing"

func TestSortAscending(t *testing.T) {
	arr := []int{2, 3, 5, 1, 4}
	Sort(arr, func(a, b int) bool {
		return a > b
	})

	res := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		if res[i] != arr[i] {
			t.Errorf("Expected result: %d. Actual result: %d", res[i], arr[i])
		}
	}
}

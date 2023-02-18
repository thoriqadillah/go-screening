package array

// Insertion sort
func Sort(numbers []int, callback func(a int, b int) bool) {
	for i := range numbers {
		for j := i; j > 0; j-- {
			if callback(numbers[j-1], numbers[j]) {
				numbers[j-1], numbers[j] = numbers[j], numbers[j-1]
			}
		}
	}
}

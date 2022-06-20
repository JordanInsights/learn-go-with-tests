package arrays_and_slices

func Sum(numbers []int) int {
	var total int
	for _, num := range numbers {
		total += num
	}

	return total
}

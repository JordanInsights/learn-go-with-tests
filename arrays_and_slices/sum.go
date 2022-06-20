package arrays_and_slices

func Sum(numbers []int) int {
	var total int
	for _, num := range numbers {
		total += num
	}

	return total
}

func SumAll(slices ...[]int) []int {
	var sliceValues []int
	for _, slice := range slices {
		num := Sum(slice)
		sliceValues = append(sliceValues, num)
	}
	return sliceValues
}

func SumAllTails(slices ...[]int) []int {
	var sliceValues []int
	for _, slice := range slices {
		num := Sum(slice)
		sliceValues = append(sliceValues, num-slice[0])
	}

	return sliceValues
}

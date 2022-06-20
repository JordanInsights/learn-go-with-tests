package ch4

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
		var num int
		if len(slice) == 0 {
			num = 0
		} else {
			num = Sum(slice[1:])
		}

		sliceValues = append(sliceValues, num)
	}

	return sliceValues
}

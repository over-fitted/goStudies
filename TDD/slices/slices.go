package slices

func Sum(input []int) int {
	var sum int
	for _, v := range input {
		sum += v
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		var tail []int
		if len(numbers) <= 1 {
			tail = []int{0}
		} else {
			tail = numbers[1:]
		}
		sums = append(sums, Sum(tail))
	}

	return sums
}

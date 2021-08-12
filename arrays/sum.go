package main

func Sum(numbers []int) int {
	sum := 0
	// range returns two values - the index and the value. The blank identifier ignores the index value
	for _, number := range numbers {
		sum += number
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

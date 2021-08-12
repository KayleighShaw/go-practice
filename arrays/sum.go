package main

func Sum(numbers [5]int) int {
	sum := 0
	// range returns two values - the index and the value. The blank identifier ignores the index value
	for _, number := range numbers {
		sum += number
	}
	return sum
}

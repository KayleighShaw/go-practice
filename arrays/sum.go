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
	lengthOfNumbers := len(numbersToSum)
	//make allows you to create a slice with a starting capacity of the len of the numbersToSum
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

package main

func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	sumOfAll := []int{}
	for _, numbers := range numbersToSum {
		sum := 0
		for _, number := range numbers {
			sum += number
		}
		sumOfAll = append(sumOfAll, sum)
	}
	return sumOfAll
}

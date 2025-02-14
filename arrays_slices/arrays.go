package main

func Sum(numbers []int) int {
	var total int
	for _, num := range numbers {
		total += num
	}
	return total
}

func SumAll(numbersToSum ...[]int) []int {
	var sum []int
	for _, number := range numbersToSum {
		sum = append(sum, Sum(number))
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sum []int
	for _, number := range numbersToSum {
		if len(number) == 0 {
			sum = append(sum, 0)
		} else {
			tails := number[1:]
			sum = append(sum, Sum(tails))

		}

	}
	return sum
}

package main

func Sum(numbers []int) int {
	var total int
	for _, num := range numbers {
		total += num
	}
	return total
}

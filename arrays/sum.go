package arrays

func Sum(numbers []int) int {

	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(mumbersToSum ...[]int) (sums []int) {

	for _, numbers := range mumbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}

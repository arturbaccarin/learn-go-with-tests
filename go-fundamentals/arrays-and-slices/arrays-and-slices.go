package arraysandslices

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

// Variadic functions can be called with any number of trailing arguments.
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)

	// make allows you to create a slice with a starting capacity of the len of the numbersToSum we need to work through.
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

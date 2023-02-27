package arrays

func Sum(arr []int) (sum int) {
	for _, number := range arr {
		sum += number
	}
	return sum
}

func SumAll(arr ...[]int) (sums []int) {
	for _, numbers := range arr {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

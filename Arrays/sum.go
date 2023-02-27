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

func SumAllTails(arr ...[]int) (sumsTails []int) {
	for _, numbers := range arr {
		if len(numbers) == 0 {
			sumsTails = append(sumsTails, 0)
		} else {
			sumsTails = append(sumsTails, Sum(numbers[1:]))
		}
	}
	return sumsTails
}

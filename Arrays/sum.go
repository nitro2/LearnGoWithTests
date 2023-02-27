package arrays

func Sum(arr []int) (sum int) {
	for i := 0; i < cap(arr); i++ {
		sum += arr[i]
	}
	return sum
}

package array

// SumAll return every summery of array
func SumAll(numbertoSum ...[]int) (sum []int) {
	//	var sum []int
	for _, number := range numbertoSum {
		sum = append(sum, Sum(number))
	}
	return
}

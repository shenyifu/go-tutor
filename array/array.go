package array

// Sum return summary of int array
func Sum(number []int) (sum int) {
	sum = 0
	for _, value := range number {
		sum += value
	}
	return
}

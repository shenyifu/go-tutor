package array

//SumAllTail return summary of tail
func SumAllTail(numbers ...[]int) (sums []int) {
	for _, numb := range numbers {
		if len(numb) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numb[1:]))
		}

	}
	return
}

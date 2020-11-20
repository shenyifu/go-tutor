package iteration

//Repeat return n times of item
func Repeat(item string, times int) string {
	result := ""
	for i := 0; i < times; i++ {
		result += item
	}
	return result
}

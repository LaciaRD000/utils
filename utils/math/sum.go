package math

func ISum(i []int) int {
	var sum int
	for _, n := range i {
		sum += n
	}
	return sum
}

package math

func Sum(i []int) int {
	var sum int
	for _, n := range i {
		sum += n
	}
	return sum
}

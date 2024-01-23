package math

func Avg(i []int) int {
	var sum int
	for _, n := range i {
		sum += n
	}
	sum = sum / len(i)
	return sum
}

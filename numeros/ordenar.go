package numeros

func SumPares(numbers []int) int {
	var sum int
	for _, n := range numbers {
		if n%2 == 0 {
			sum += n
		}
	}
	return sum
}

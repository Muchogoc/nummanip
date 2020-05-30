package calc

// Add is used to return sum numbers
func Add(numbers ...int) int {
	sum := 0

	for _, number := range numbers {
		sum = sum + number
	}

	return sum
}

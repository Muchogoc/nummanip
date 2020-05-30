package calc

import "errors"

// Add is used to return sum numbers
func Add(numbers ...int) (int, error) {
	sum := 0

	if len(numbers) < 2 {
		return sum, errors.New("provide more than two numbers")
	}
	for _, number := range numbers {
		sum = sum + number
	}

	return sum, nil
}

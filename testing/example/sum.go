package example

import (
	"errors"
)

func Suma(a, b int) int {
	return a + b
}

func Subtracao(a, b int) (int, error) {
	if a < b {
		return 0, errors.New("error in subtracao")
	}

	return a - b, nil
}

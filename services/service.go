package services

import (
	. "strconv"
)

func Add(a, b int) int {
	return a + b
}

func Sub(a, b string) (int, error) {
	resa, err := Atoi(a)
	if err != nil {
		return 0, err
	}
	resb, err := Atoi(b)
	if err != nil {
		return 0, err
	}
	return resa - resb, nil
}

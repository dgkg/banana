package chainfunc

import (
	"errors"
	"math"
)

var (
	ErrorOutOfBoundaries   = errors.New("out ouf boundaries")
	ErrorOutOfDivideByZero = errors.New("forbidden to divide by zero")
)

type myInt int

func (i myInt) Add(j int) (myInt, error) {
	res := int64(i) + int64(j)
	if int64IsOutOfBoundaries(res) {
		return 0, ErrorOutOfBoundaries
	}
	return myInt(res), nil
}

func (i myInt) Sub(j int) (myInt, error) {
	res := int64(i) - int64(j)
	if int64IsOutOfBoundaries(res) {
		return 0, ErrorOutOfBoundaries
	}
	return myInt(res), nil
}

func (i myInt) Divide(j int) (myInt, error) {
	if j == 0 {
		return 0, ErrorOutOfDivideByZero
	}
	return myInt(int64(i) / int64(j)), nil
}

func (i myInt) Multiply(j int) (myInt, error) {
	if j == 0 {
		return 0, nil
	}
	res := int64(i) * int64(j)
	if int64IsOutOfBoundaries(res) {
		return 0, ErrorOutOfBoundaries
	}
	return myInt(res), nil
}

func intIsOutOfBoundaries(i int) bool {
	return i < math.MinInt32 || i > math.MaxInt32
}

func int64IsOutOfBoundaries(i int64) bool {
	return i < math.MinInt32 || i > math.MaxInt32
}

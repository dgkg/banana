package chainfunc

import (
	"math"
	"testing"
)

func TestMyIntAdd(t *testing.T) {
	data := []struct {
		start    myInt
		param    int
		expected myInt
		err      error
	}{
		{1, 2, 3, nil},
		{1, math.MaxInt64, 0, ErrorOutOfBoundaries},
		{math.MaxInt32, math.MaxInt32, 0, ErrorOutOfBoundaries},
		{math.MinInt32, math.MaxInt64, 0, ErrorOutOfBoundaries},
	}
	for _, v := range data {
		var my myInt = myInt(v.start)
		res, err := my.Add(v.param)
		if res != v.expected {
			t.Errorf("Expected %d but got %d", v.expected, res)
		}
		if err != v.err {
			t.Errorf("Expected %v but got %v", v.err, err)
		}
	}
}

func TestMyIntSub(t *testing.T) {

}

func TestMyIntDivide(t *testing.T) {

}

func TestMyIntMultiply(t *testing.T) {

}

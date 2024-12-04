package chainfunc

import (
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
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

func FuzzMyIntAddSub(f *testing.F) {
	testparam := []int{1, math.MaxInt64, math.MinInt64}
	for _, tc := range testparam {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	var i myInt = 0
	f.Fuzz(func(t *testing.T, orig int) {
		i, err := i.Add(orig)
		if err != nil {
			return
		}
		i, err = i.Sub(int(i))
		if err != nil {
			return
		}
		if i != 0 {
			t.Errorf("Before: %v, after: %v", 0, i)
		}
	})
}

func TestWithConvey(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		var x myInt = 1

		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 3)
			})
		})
	})
}

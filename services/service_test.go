package services_test

import (
	// stdlib
	"testing"
	// contribs

	// local
	"github.com/dgkg/banana/services"
)

func TestAdd(t *testing.T) {
	data := []struct {
		a, b, expected int
	}{
		{1, 2, 3},
	}
	for _, v := range data {
		res := services.Add(v.a, v.b)
		if res != v.expected {
			t.Errorf("Expected %d but got %d", v.expected, res)
		}
	}
}

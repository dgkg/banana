package slices

import (
	"reflect"
	"testing"
)

func isIntricated(a, b []int) bool {
	if len(a) == 0 || len(b) == 0 {
		return true
	}
	a[0] = 1234567890
	if !reflect.DeepEqual(a, b) {
		return false
	}
	return true
}

func TestCopiedSlice(t *testing.T) {
	tbl := []int{1, 2, 3, 4, 5, 6}
	res := CopiedSlice(tbl)

	if isIntricated(tbl, res) {
		t.Error("Expected", tbl, "got", res)
	}
}

func TestAppendedSlice(t *testing.T) {
	tbl := []int{1, 2, 3, 4, 5, 6}
	res := AppendedSlice(tbl)
	if isIntricated(tbl, res) {
		t.Error("Expected", tbl, "got", res)
	}
}
func TestAppendedSlice2(t *testing.T) {
	tbl := []int{1, 2, 3, 4, 5, 6}
	res := AppendedSlice2(tbl)
	if isIntricated(tbl, res) {
		t.Error("Expected", tbl, "got", res)
	}
}

func TestAppendedSlice3(t *testing.T) {
	tbl := []int{1, 2, 3, 4, 5, 6}
	res := AppendedSlice3(tbl)
	if isIntricated(tbl, res) {
		t.Error("Expected", tbl, "got", res)
	}
}

func TestCloneSlice(t *testing.T) {
	tbl := []int{1, 2, 3, 4, 5, 6}
	res := CloneSlice(tbl)
	if !isIntricated(tbl, res) {
		t.Error("Expected", tbl, "got", res)
	}
}
func TestCloneSlice2(t *testing.T) {
	tbl := []int{1, 2, 3, 4, 5, 6}
	res := CloneSlice2(tbl)
	if !isIntricated(tbl, res) {
		t.Error("Expected", tbl, "got", res)
	}
}

func TestAppendCapacity(t *testing.T) {
	for i := 0; i < 5000; i++ {
		AppendCapacity(i)
	}
}

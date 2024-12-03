package slices

func CopiedSlice(s []int) []int {
	copied := make([]int, len(s))
	copy(copied, s)
	return copied
}

func AppendedSlice(s []int) []int {
	appended := append([]int(nil), s...)
	return appended
}

func AppendedSlice2(s []int) []int {
	appended := make([]int, 0, len(s))
	appended = append(appended, s...)
	return appended
}

func AppendedSlice3(s []int) []int {
	appended := make([]int, 0, len(s))
	appended = append(appended, s[:]...)
	return appended
}

func CloneSlice(s []int) []int {
	cloned := s
	return cloned
}

func CloneSlice2(s []int) []int {
	cloned := s[:]
	return cloned
}

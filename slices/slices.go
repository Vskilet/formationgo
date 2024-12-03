package slices

func Copy(a []int) []int{
	b := make([]int, len(a))
	copy(b, a)
	b[1] = 77
	return b
}
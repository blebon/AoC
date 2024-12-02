package util

// RemoveElementInSlice returns a new slice with the element
// at index i removed, without modifying the original slice.
func RemoveElementInSlice(slice []int, i int) []int {
	c := make([]int, len(slice))
	for i, v := range slice {
		c[i] = v
	}
	s := append(c[:i], c[i+1:]...)
	return s
}

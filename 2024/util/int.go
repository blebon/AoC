package util

const MAX_INT = int(^uint(0) >> 1)

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

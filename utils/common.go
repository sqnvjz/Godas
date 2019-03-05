package utils

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func InArray(i int, arr []int) bool {
	for _, v := range arr {
		if i == v {
			return true
		}
	}
	return false
}

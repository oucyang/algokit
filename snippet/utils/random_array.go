package utils

import "math/rand"

func RandomIntArray(maxSize int, maxValue int) []int {
	if maxValue < 0 {
		maxValue = -maxValue
	}
	double := maxValue * 2
	arrLen := rand.Intn(maxSize + 1)
	var arr = make([]int, arrLen)
	for i := 0; i < arrLen; i++ {
		arr[i] = rand.Intn(double) - maxValue
	}
	return arr
}

func IsSameIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	var m = make(map[int]int, 0)
	for _, n := range a {
		m[n] = m[n] + 1
	}
	for _, n := range b {
		m[n] = m[n] - 1
	}
	for _, c := range m {
		if c != 0 {
			return false
		}
	}
	return true
}

func CopyIntSlice(ints []int) []int {
	replica := make([]int, len(ints))
	copy(replica, ints)
	return replica
}

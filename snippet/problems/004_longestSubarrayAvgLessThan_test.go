package problems

import (
	"github.com/oucyang/algokit/snippet/utils"
	"testing"
)

// O(n^2)
func LongestSubarrayAvgLessThanSlowWay(arr []int, v int) int {
	if len(arr) < 1 {
		return 0
	}
	ans := 0
	for i := 0; i < len(arr); i++ {
		sum, length := 0, 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			length = j - i + 1
			if sum <= v*length && ans < length {
				ans = length
			}
		}
	}
	return ans
}

type LongestSubarrayAvgLessThanCase struct {
	Arr []int
	V   int
	Ans int
}

func randomLongestSubarrayAvgLessThanCase(maxSize, maxValue int) *LongestSubarrayAvgLessThanCase {
	var arr = utils.RandomIntArray(maxSize, maxValue)
	var v = utils.RandomInt(maxValue)
	return &LongestSubarrayAvgLessThanCase{
		Arr: arr,
		V:   v,
		Ans: LongestSubarrayAvgLessThanSlowWay(arr, v),
	}
}

func TestLongestSubarrayAvgLessThan(t *testing.T) {
	N := 500000
	MaxSize := 100
	MaxValue := 100000

	var testOneCase = func(c *LongestSubarrayAvgLessThanCase) {
		if ans := LongestSubarrayAvgLessThan(c.Arr, c.V); ans != c.Ans {
			t.Fatalf("wrong answer %d case is %+v", ans, *c)
		}
	}

	for i := 0; i < N; i++ {
		testOneCase(randomLongestSubarrayAvgLessThanCase(MaxSize, MaxValue))
	}
}

package sort

import (
	"math/rand"
	"sort"
)

func QuickSortV1(objs sort.Interface) {
	if objs == nil || objs.Len() < 2 {
		return
	}
	quickSort1(objs, 0, objs.Len()-1)
}

func quickSort1(objs sort.Interface, l, r int) {
	if l >= r {
		return
	}
	var m = partition1(objs, l, r)
	quickSort1(objs, l, m-1)
	quickSort1(objs, m+1, r)
}

func QuickSortV2(nums []int) {
	if len(nums) < 2 {
		return
	}
	quickSort2(nums, 0, len(nums)-1)
}

func quickSort2(nums []int, l, r int) {
	if l >= r {
		return
	}
	less, more := partition2(nums, l, r)
	quickSort2(nums, l, less-1)
	quickSort2(nums, more+1, r)
}

func QuickSortV3(nums []int) {
	if len(nums) < 2 {
		return
	}
	quickSort3(nums, 0, len(nums)-1)
}

func quickSort3(nums []int, l, r int) {
	if l >= r {
		return
	}
	randIndex := rand.Intn(r-l) + l
	nums[randIndex], nums[r] = nums[r], nums[randIndex]
	less, more := partition2(nums, l, r)
	quickSort3(nums, l, less-1)
	quickSort3(nums, more+1, r)
}

// partition by objs[r], <objs[r], >=objs[r]
func partition1(objs sort.Interface, l, r int) int {
	if l > r {
		return -1
	}
	if l == r {
		return l
	}
	var lessEqual, index = l, l
	for index < r {
		if objs.Less(index, r) == true {
			objs.Swap(index, lessEqual)
			lessEqual++
		}
		index++
	}
	objs.Swap(lessEqual, r)
	return lessEqual
}

// partition by objs[r], <objs[r], ==objs[r], >objs[r]
func partition2(nums []int, l, r int) (int, int) {
	if l > r {
		return -1, -1
	}
	if l == r {
		return l, r
	}
	var less, more, index = l, r, l
	for index < more {
		if nums[index] == nums[r] {
			index++
		} else if nums[index] < nums[r] {
			nums[index], nums[less] = nums[less], nums[index]
			less++
			index++
		} else {
			more--
			nums[index], nums[more] = nums[more], nums[index]
		}
	}
	nums[more], nums[r] = nums[r], nums[more]
	return less, more
}

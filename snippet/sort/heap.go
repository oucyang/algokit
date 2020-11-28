package sort

func heapify(nums []int, i, stop int) {
	child := (i << 1) + 1
	for child < stop {
		if right := child + 1; right < stop && nums[right] > nums[child] {
			child = right
		}
		if nums[child] > nums[i] {
			nums[i], nums[child] = nums[child], nums[i]
			i = child
		} else {
			return
		}
		child = (i << 1) + 1
	}
}

func HeapSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	for i := len(nums) - 1; i > -1; i-- {
		heapify(nums, i, len(nums))
	}
	for last := len(nums) - 1; last > 0; last-- {
		nums[0], nums[last] = nums[last], nums[0]
		heapify(nums, 0, last)
	}
}

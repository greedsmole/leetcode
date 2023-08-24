func search(nums []int, target int) int {
	n := len(nums)
	if n == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	if nums[0] < nums[n-1] {
		return bs2(nums, 0, n-1, target)
	}
	minIndex := bs(nums, 0, n-1, nums[n-1])
	maxIndex := minIndex - 1

	if nums[n-1] >= target && nums[minIndex] <= target {
		return bs2(nums, minIndex, n-1, target)
	} else if nums[0] <= target && nums[maxIndex] >= target {
		return bs2(nums, 0, maxIndex, target)
	}
	return -1
}

func bs2(nums []int, l int, r int, c int) int {
	mid := (l + r) / 2
	if nums[mid] == c {
		return mid
	}
	if l == r {
		return -1
	}
	if nums[mid] > c {
		return bs2(nums, l, mid, c)
	} else if nums[mid] < c {
		return bs2(nums, mid+1, r, c)
	}
	return -1
}

func bs(nums []int, a int, b int, c int) int {
	n := (a + b) / 2
	if a >= b {
		return a
	}
	if nums[n] < c {
		b = n
		c = nums[n]
	} else if nums[n] > c {
		a = n + 1
	}
	return bs(nums, a, b, c)
}
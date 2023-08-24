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

func findMin(nums []int) int {
	n := len(nums)
	if nums[0] < nums[n-1] {
		return nums[0]
	}
	nmin := bs(nums, 0, n-1, nums[n-1])
	return nums[nmin]
}
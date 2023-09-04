package main

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func robrec(nums []int, i int) int {
	if i < 0 {
		return 0
	}
	if cache[i] >= 0 {
		return cache[i]
	}
	res := max(robrec(nums, i-2)+nums[i], robrec(nums, i-1))
	cache[i] = res
	return res
}

var cache []int

func rob(nums []int) int {
	cache = make([]int, len(nums)+1)
	for i := 0; i < len(cache); i++ {
		cache[i] = -1
	}
	return robrec(nums, len(nums)-1)
}

package main

import "fmt"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	h := make(map[int]bool)
	for _, v := range nums {
		h[v] = true
	}
	best := 0
	for x := range nums {
		if !h[x-1] {
			y := x + 1
			for h[y] {
				y++
			}
			best = max(best, y-x)
		}
	}
	return best
}

func main() {
	nums := []int{
		-9, -8, -7, -4, -4, -4, -3, -2, -1, 0, 1, 1, 3, 5, 6, 7}
	fmt.Println(longestConsecutive(nums))
}

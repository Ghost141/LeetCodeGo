package main

func removeElement(nums []int, val int) int {
	ind := 0
	for i, v := range nums {
		if val != v {
			nums[ind] = nums[i]
			ind++
		}
	}
	return ind
}

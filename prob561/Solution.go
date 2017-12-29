package main

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i, num := range nums {
		if i % 2 == 0 {
			res += num
		}
	}
	return res
}

func main() {
	print(arrayPairSum([]int{1, 2, 3}))
}

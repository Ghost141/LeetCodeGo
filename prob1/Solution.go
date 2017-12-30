package main

import "fmt"

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	m := make(map[int]int)
	for i, num := range nums {
		if j, ok := m[target - num]; ok {
			res[0] = j
			res[1] = i
			return res
		} else {
			m[num] = i
		}
	}

	return res
}

func main() {
	fmt.Println("Hello world")
	fmt.Println(twoSum([]int {3,2,4}, 6))
}
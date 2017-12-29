package main

import "fmt"

func matrixReshape(nums [][]int, r int, c int) [][]int {
	n, m := len(nums), len(nums[0])
	if n*m != r*c {
		return nums
	}

	res := make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
		for j := range res[i] {
			oi, oj := transfer(i, j, m, c)
			res[i][j] = nums[oi][oj]
		}
	}
	return res
}

func transfer(i, j, m, c int) (newI, newJ int) {

	newI = int((i*c + j) / m)
	newJ = (i*c + j) % m

	return
}

func main() {
	fmt.Println(matrixReshape([][]int{
		{1, 2},
		{3, 4},
	}, 4, 1))
	fmt.Println(matrixReshape([][]int{
		{1, 2},
		{3, 4},
	}, 1, 4))
}

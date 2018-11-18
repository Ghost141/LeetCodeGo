package main

import "strconv"

func convertToBase7(num int) string {
	return strconv.FormatInt(int64(num), 7)
}

func main() {
	println(convertToBase7(1))
	println(convertToBase7(5))
	println(convertToBase7(7))
	println(convertToBase7(100))
}

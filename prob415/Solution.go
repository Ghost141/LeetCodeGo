package main

import (
	"math"
	"strconv"
	"strings"
)

func addStrings(num1 string, num2 string) string {
	var s1, s2, ind, c int64
	maxLen := int64(math.Max(float64(len(num1)), float64(len(num2))))
	res := make([]byte, maxLen)
	i := len(num1) - 1
	j := len(num2) - 1
	ind = int64(math.Max(float64(i), float64(j)))

	for i >= 0 || j >= 0 {
		if i >= 0 {
			s1 = int64([]byte(num1)[i] - '0')
		} else {
			s1 = 0
		}
		if j >= 0 {
			s2 = int64([]byte(num2)[j] - '0')
		} else {
			s2 = 0
		}

		sum := s1 + s2 + c
		res[ind] = byte(sum%10) + '0'
		c = sum / 10

		ind--
		i--
		j--
	}

	if c > 0 {
		return strings.Join([]string{strconv.FormatInt(c, 10), string(res)}, "")
	}

	return string(res)
}

func main() {
	println(addStrings("123", "123"))
	println(addStrings("1", "123"))
	println(addStrings("9", "11"))
}

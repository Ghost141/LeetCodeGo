package main

import "strings"

func intToRoman(num int) string {
	dict := [][]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "M", "MM", "MMM"},
	}

	res := make([]string, 0, 4)
	res = append(res, dict[3][(num/1000)%10])
	res = append(res, dict[2][(num/100)%10])
	res = append(res, dict[1][(num/10)%10])
	res = append(res, dict[0][num%10])

	return strings.Join(res, "")
}

func main() {
	println(intToRoman(4))
}

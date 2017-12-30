package main

import (
	"fmt"
	"strings"
)

var keyboard = map[int32]int{
	'a': 1,
	'b': 2,
	'c': 2,
	'd': 1,
	'e': 0,
	'f': 1,
	'g': 1,
	'h': 1,
	'i': 0,
	'j': 1,
	'k': 1,
	'l': 1,
	'm': 2,
	'n': 2,
	'o': 0,
	'p': 0,
	'q': 0,
	'r': 0,
	's': 1,
	't': 0,
	'u': 0,
	'v': 2,
	'w': 0,
	'x': 2,
	'y': 0,
	'z': 2,
}

func findWords(words []string) (res []string) {
	for i, str := range words {
		str = strings.ToLower(str)
		row := keyboard[int32(str[0])]
		add := true
		for _, c := range str {
			if keyboard[c] != row {
				add = false
				break
			}
		}
		if add {
			res = append(res, words[i])
		}
	}
	return
}

func main() {
	fmt.Println(findWords([]string {"Hello"}))
}
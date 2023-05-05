package main

import (
	"fmt"
)

var iniSlices = []int8{123, 100, 90}

func SwapRune(r rune) rune {
	switch {
		case 97 <= r && r <= 122:
			return r - 32
		case 65 <= r && r <= 90:
			return r + 32
		default:
			return r
	}
}

func main() {
	fmt.Println(SwapRune('a'))
}



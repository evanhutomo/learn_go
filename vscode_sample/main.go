package main

import (
	"fmt"

	"eh.org/vscode_sample/mascot"
	"rsc.io/quote"
)

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func main() {
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Go())

	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
}

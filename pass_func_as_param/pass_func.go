package main

import "fmt"

type fn func(int)
type fn_2arg func(int, string)

func myfn1(i int) {
	fmt.Printf("i is %v\n", i)
}

func myfn2(i int) {
	fmt.Printf("i is %v\n", i)
}

func fn2arg(angka int, tulisan string) {
	fmt.Printf("angka: %d, string: %s\n", angka, tulisan)
}

func test(f fn, val int) {
	f(val)
}

func test2(f2 fn_2arg, angka int, tulisan string) {
	f2(angka, tulisan)
}

func main() {
	test(myfn1, 123)
	test(myfn2, 4567)
	test2(fn2arg, 111, "yeah")
}

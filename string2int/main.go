package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	strVar := "56"
	intVar, err := strconv.Atoi(strVar)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))
}

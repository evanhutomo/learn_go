package main

import (
	"github.cim/gosuri/uiprogress"
)

func main() {
	uiprogress.Start()
	bar := uiprogress.AddBar(100)
	bar.AppendCompleted()

}

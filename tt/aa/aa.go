package main

import (
	"time"

	"github.com/gosuri/uiprogress"
)

func main() {
	uiprogress.Start()
	bar := uiprogress.AddBar(100)

	// optionally, append and prepend completion and elapsed time
	bar.AppendCompleted()
	bar.PrependElapsed()

	for bar.Incr() {
		time.Sleep(time.Millisecond * 20)
	}
}

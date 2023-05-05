package main

import "fmt"

type reader interface {
	read() (int, error)
}

type file struct {
	name        string
	fileContent string
}

func (f *file) read() (int, error) {
	content := "file contents"
	f.fileContent = content
	return len(content), nil
}

type pipe struct {
	name        string
	pipeMessage string
}

func (p *pipe) read() (int, error) {
	msg := `{name: "Henry", title: "Developer"}`
	p.pipeMessage = msg
	return len(msg), nil
}

// implement reader interface
func retrieve(r reader) error {
	len, _ := r.read()
	fmt.Println(fmt.Sprintf("read %d bytes", len))
	return nil
}

func main() {
	f := file{
		name: "data.json",
	}
	p := pipe{
		name: "pipe_message",
	}

	retrieve(&f)
	fmt.Println(f.fileContent)
	retrieve(&p)
	fmt.Println(p.pipeMessage)
}

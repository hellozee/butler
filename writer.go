package main

import (
	"fmt"
	"os"
)

type Writer struct {
	FileName string
	data     string
}

func NewWriter(fname string) *Writer {
	w := Writer{
		FileName: fname,
		data:     "",
	}
	return &w
}

func (w *Writer) PrintCommand(cmd string) {
	w.data += "```bash\n"
	w.data += "bash:- $ " + cmd + "\n"
	fmt.Println(cmd)
}

func (w *Writer) PrintOutput(stdout, stderr string) {
	w.data += stdout + stderr
	w.data += "```\n"
}

func (w *Writer) Save() {
	file, err := os.Create(w.FileName)
	must(err)
	defer file.Close()

	fmt.Fprintf(file, w.data)
}

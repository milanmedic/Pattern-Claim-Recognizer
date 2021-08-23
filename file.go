package main

import (
	"os"
)

type File struct {
	name     string
	contents []byte
}

func (f File) OpenFile() {
	_, err := os.Open(f.name)
	if err != nil {
		HandleError(err, "terminate")
	}
}

func (f File) ReadFile() string {
	contents, err := os.ReadFile(f.name)
	f.contents = contents
	if err != nil {
		HandleError(err, "terminate")
	}
	return ByteArrToStr(f.contents)
}

package main

import (
	"io"
	"os"
	"strings"
)

// Using only the provided packages io, os and strings.

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	size, err := rot.r.Read(b)
	for i := 0; i < size; i++ {
		val := b[i]
		b[i] = rot13(val)
	}
	return size, err
}

func rot13(char byte) byte {
	lookup := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	index := strings.Index(lookup, strings.ToUpper(string(char)))
	if index < 0 {
		return char
	}
	rotIndex := (13 + index) % 26

	match := lookup[rotIndex]

	//Check if char is a lower case then convert the macth back to lower case.
	if string(char) != string(lookup[index]) {
		return strings.ToLower(string(match))[0]
	} else {
		return match
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

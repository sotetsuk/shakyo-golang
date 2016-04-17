package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func isAlphabet(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func (rot *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot.r.Read(b)
	if err != nil {
		return n, err
	}

	for i, v := range b {
		if !isAlphabet(v) {
			continue
		} else if r := v + 13; isAlphabet(r) {
			b[i] = r
		} else if r := v - 13; isAlphabet(r) {
			b[i] = r
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

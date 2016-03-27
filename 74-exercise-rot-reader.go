package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (n int, err error){
	n, err = rot.r.Read(b)
	for i, v :=  range b {
		if r := v + 13; (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			b[i] = r
		} else if r := v - 13; (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			b[i] = r
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

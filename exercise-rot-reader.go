package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	if err != nil {
		return n, err
	}
	for i := 0; i < n; i++ {
		base := byte(0)
		if b[i] >= 'a' && b[i] <= 'z' {
			base = 'a'
		} else if b[i] >= 'A' && b[i] <= 'Z' {
			base = 'A'
		}
		if base != 0 {
			b[i] = (b[i]+13-base)%26 + base
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

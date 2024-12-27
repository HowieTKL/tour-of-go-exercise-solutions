package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(b []byte) (int, error) {
	n, err := rr.r.Read(b)
	if (err != nil) {
		return n, err
	}
	for x := 0; x < len(b); x++ {
		if b[x] >= 65 && b[x] <= 77 {
			b[x] += 13
		} else if b[x] >=78 && b[x] <= 90 {
			b[x] -= 13
		} else if b[x] >= 97 && b[x] <= 109 {
			b[x] += 13
		} else if b[x] >=110 && b[x] <= 122 {
			b[x] -= 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

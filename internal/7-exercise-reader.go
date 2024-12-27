package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (MyReader) Read(b []byte) (int, error) {
    for x := range b {
        b[x] = 'A'
    }
    return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}

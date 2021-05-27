package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader. that generates infinte 'A'

func (reader MyReader) Read(b []byte) (int, error) {
	ret := 0
	for index := range b {
		b[index] = 'A'
		ret++
	}
	return ret, nil
}

func main() {
	reader.Validate(MyReader{})
}

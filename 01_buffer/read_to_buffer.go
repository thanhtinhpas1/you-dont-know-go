package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := bytes.NewBuffer([]byte("hello read from buffer"))
	b := make([]byte, 5)

	n, err := buf.Read(b)
	if err != nil {
		panic("error read from buffer")
	}

	fmt.Println(string(b[:n]))
}

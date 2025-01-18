package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := bytes.NewBuffer([]byte("hello write "))
	buf.Write([]byte("buffer!"))

	fmt.Println(buf.String())
}

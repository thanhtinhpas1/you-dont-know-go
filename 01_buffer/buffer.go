package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf1 := bytes.NewBuffer([]byte("hello, buffer!"))
	buf2 := bytes.NewBufferString("hello, buffer string!")
	fmt.Println(buf1.String())
	fmt.Println(buf2.String())
}

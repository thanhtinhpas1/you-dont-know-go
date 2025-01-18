package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := bytes.NewBuffer([]byte("hello buffer"))
	fmt.Println(buf.Len())
	fmt.Println(buf.Cap())
	fmt.Println(buf.String())
}

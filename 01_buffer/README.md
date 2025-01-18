# How to use Buffer in Go

## What is Buffer?
A buffer is a temporary storage area that was using to hold data while data is being moved from one place to another place.
?
Buffers are useful for several reasons:
1. *Efficient data manipulation*: Buffer allow you to efficiently manipulate data in memory befopre processing & transmitting it
2. *Performance*: Using buffer can improve the performance of I/O operations by reducing the number of system calls
3. *Convenience*: Buffer provide a convenieent way to build and modify byte slices.

## Working with buffer in Go?
The `bytes` package in Go provide the `Buffer` type, which is a dynamic buffer of bytes. Let's explore how to use Buffer with some practical examples:

### Creating a Buffer
To Create a new Buffer, you can use the `bytes.NewBuffer` or `bytes.NewBufferString` functions. Here's how:

```go
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


// Output
// > go run buffer.go
// hello, buffer!
// hello, buffer string!
```

## Writing to a Bufer
You can write data to a buffer using the `Write` method. This method appends the given bytes to the buffer.

```go
import (
	"bytes"
	"fmt"
)

func main() {
	buf := bytes.NewBuffer([]byte("hello write "))
	buf.Write([]byte("buffer!"))

	fmt.Println(buf.String())
}


//> go run write_to_buffer.go
//hello write buffer!
```

## Reading from a Buffer
Reading data from a buffer can be done using the Read method, which reads up to `len(p)` bytes into p.

```go
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

//> go run read_to_buffer.go
//hello
```

## Other useful methods
The Buffer type provides several other useful methods, such as:
- `WriteString(s string)`: writes a string to the buffer
- `Bytes() []byte`: returns the buffer's content as a byte slice
- `String() string`: returns the buffer's content as a string
- `Len() int`: returns the number of bytes currently in the buffer
- `Cap() int`: returns the buffer's capacity

```go
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

//> go run useful_methods.go
//12
//12
hello buffer
```

## Create Reader with buffer
We also can create Reader with the specified buffer size, and having buffer for scanning when reading lines in the file for example

```go

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fp := newFileParser()
	const maxCapacity = 1024 * 1024
	buf := make([]byte, maxCapacity)

	reader := bufio.NewReaderSize(file, maxCapacity)

	scanner := bufio.NewScanner(reader)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		line := scanner.Bytes()
		fp.parseLine(line)
	}
}

```
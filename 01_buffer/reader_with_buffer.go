package main

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

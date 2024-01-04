package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buffer := make([]byte, 1024)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		panic(err)
	}
	data := buffer[:n]
	print(string(data))
}

package main

import (
	"fmt"
	"os"
)

func main() {
	argWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	argsWithoutProg1 := os.Args[2:]
	arg := os.Args[3]
	fmt.Println(argWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(argsWithoutProg1)
	fmt.Println(arg)
}

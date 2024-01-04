package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestWaitgroup(t *testing.T) {
	WaitGroupExample()
}

func WaitGroupExample() {

	fmt.Println("START --------")

	var wg sync.WaitGroup
	wg.Add(4)

	go Hello(&wg, "Vin 1")
	go HelloWorldDuy(&wg)
	go Hello(&wg, "Vin 2")
	go HelloWorldThien(&wg)

	wg.Wait()
	fmt.Println("END --------")
}

func Hello(wg *sync.WaitGroup, name string) {
	fmt.Println("hello " + name)
}

func HelloWorldDuy(wg *sync.WaitGroup) {
	wg.Done()
	fmt.Println("hello Duy")
}

func HelloWorldThien(wg *sync.WaitGroup) {
	wg.Done()
	fmt.Println("hello Thien")
}

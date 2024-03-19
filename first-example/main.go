package main

import (
	"fmt"
	"time"
)

func main() {

	go printSomething("this is the first!")

	time.Sleep(1 * time.Second)

	printSomething("this is the second!")
}

func printSomething(s string) {
	fmt.Println(s)
}

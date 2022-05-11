package main

import "fmt"

func main() {
	println("Hello, world.")
}

func println(args ...interface{}) {
	fmt.Println(args...)
}

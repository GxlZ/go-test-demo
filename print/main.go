package main

import "fmt"

const helloPrefix = "Hello, "

func hello(name string) string {
	if name == ""{
		name = "world"
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(hello("go-kit"))
}

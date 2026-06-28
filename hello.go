package main

import "fmt"

func greet() {
	fmt.Println(Hello("Dev"))
}

func Hello(name string) string {
	return "Hello, " + name
}

package main

import "fmt"

const someNumber = 10

// simple function to add a number
func adder(num int) int {
	return num + 1
}

func main() {
	result := adder(someNumber)
	fmt.Printf("hello %d + 1= %d\n", someNumber, result)
}

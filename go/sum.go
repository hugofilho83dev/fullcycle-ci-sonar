package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2))
}

func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

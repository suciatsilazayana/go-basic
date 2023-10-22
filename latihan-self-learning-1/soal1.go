package main

import (
	"fmt"
)

func Fibonacci(n int) {

	var n3, n1, n2 int = 0, 1, 1

	for i := 1; i <= n; i++ {

		fmt.Print(n1, ", ")

		n3 = n1 + n2

		n1 = n2

		n2 = n3

	}
}

func main() {
	n := 10
	Fibonacci(n)
}

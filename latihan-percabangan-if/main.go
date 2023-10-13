package main

import "fmt"

func main() {
	num := 3

	if num%2 == 0 {
		fmt.Printf("%d merupakan bilangan genap", num)
	} else {
		fmt.Printf("%d merupakan bilangan ganjil", num)
	}
}

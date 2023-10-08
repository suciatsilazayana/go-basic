package main

import "fmt"

func main() {
	var huruf string = "i"

	switch huruf {
	case "a", "i", "u", "e", "o":
		fmt.Printf("Huruf %v adalah huruf vokal", huruf)
	default:
		fmt.Printf("Huruf %v adalah huruf konsonan", huruf)
	}
}

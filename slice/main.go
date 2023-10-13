package main

import "fmt"

func main() {
	iniArray1 := [4]string{"Cat", "Dog", "Bird"}
	iniArray2 := [...]string{"Honda", "Yamaha", "Mitsubishi"}
	iniSlice := []string{"Apple", "Mango", "Avocado"}

	fmt.Println(iniArray1)
	fmt.Println(iniArray2)
	fmt.Println(iniSlice)
}

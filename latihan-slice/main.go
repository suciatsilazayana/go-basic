package main

import "fmt"

func main() {
	// Diberikan sebuah array seperti berikut :
	animals := [...]string{"Cat", "Dog", "Pinguin", "Chicken", "Snake"}

	// Lalu, lengkapi variable variable berikut sesuai dengan expected-nya :
	mammals := animals[0:2]   // expected : {Cat, Dog}
	notMammals := animals[2:] // expected : {Pinguin, Chicken, Snake}
	haveLegs := animals[:4]   // expected : {Cat, Dog, Pinguin, Chicken}

	fmt.Println(mammals)
	fmt.Println(notMammals)
	fmt.Println(haveLegs)
	// Setelah itu, lakukan hal berikut :
	// 1. Ubah value Dog menjadi Cow
	mammals[1] = "Cow"
	// 2. Ubah value Pinguin menjadi Bird
	notMammals[0] = "Bird"
	haveLegs[2] = "Bird"

	fmt.Println(mammals)
	fmt.Println(notMammals)
	fmt.Println(haveLegs)
}

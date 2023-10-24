package main

import (
	"fmt"
)

func divide(num1, num2 int) int {

	if num2 == 0 {
		panic("Tidak dapat dibagi dengan 0")

	}
	return num1 / num2
}

// Recover
func catchError() {
	err := recover()
	if err != nil {
		fmt.Println("Error: Panic :", err)
	}
}

func main() {

	var num1, num2 int
	fmt.Println("Masukkan angka pertama: ")
	//mengambil inputan user
	fmt.Scan(&num1)

	fmt.Println("masukkan angka kedua: ")
	//mengambil inputan user
	fmt.Scan(&num2)

	defer catchError()

	err := recover()
	result := divide(num1, num2)
	if err != nil {
		fmt.Println("Kesalahan:", err)
	} else {
		fmt.Println("Hasil pembagian: ", result)
	}
}

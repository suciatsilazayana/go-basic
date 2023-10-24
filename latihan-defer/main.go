package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Program selesai")

	var num1, num2 int
	fmt.Println("Masukkan angka pertama: ")
	fmt.Scan(&num1)

	fmt.Println("masukkan angka kedua: ")
	fmt.Scan(&num2)

	hasil := num1 + num2

	fmt.Println("Hasilnya penjumlahan ", hasil)

}

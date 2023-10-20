package main

import (
	"fmt"
	"latihan-package-access-modifier/calculator"
)

func main() {

	hasil := calculator.Add(5, 5)
	fmt.Println("Hasil penjumlahan", hasil)
}

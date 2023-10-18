package main

import "fmt"

func main() {
	fruits := map[string]int{
		"Apple":  10,
		"Banana": 15,
		"Orange": 8,
		"Grape":  20,
	}

	fmt.Println("Sebelum ditambah/hapus")
	fmt.Println(fruits)

	// menambahkan map
	fruits["Mango"] = 12
	fruits["Strawberry"] = 7

	fmt.Println("Setelah menambahkan Mango dan Strawbery")
	fmt.Println(fruits)

	// menghapus buah orange
	delete(fruits, "Orange")

	fmt.Println("Setelah menghapus Orange")
	fmt.Println(fruits)

	// mencari item
	key := "Apple"
	val, isExist := fruits[key]

	if isExist {
		fmt.Println(key, "Jumlah Apel yang tersedia adalah", val)
	} else {
		fmt.Println(key, "Buah apel tidak ditemukan")

	}

	fmt.Println("Daftar buah-buahan beserta jumlahnya:")
	for key, value := range fruits {
		fmt.Println(key, ":", value)

	}

}

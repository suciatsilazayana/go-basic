package main

func main() {
	fruits := map[string]int{
		"Apple":  10,
		"Banana": 15,
		"Orange": 8,
		"Grape":  20,
	}

	// menambahkan map
	fruits["Mango"] = 12
	fruits["Strawberry"] = 7

	// menghapus buah orange
	delete(fruits, "Orange")
}

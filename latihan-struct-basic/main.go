package main

import "fmt"

// Membuat struct
type Book struct {
	Title  string
	Author string
}

func main() {
	// nomor 1
	var book1 Book
	book1.Title = "Pemrograman Go"
	book1.Author = "John Doe"

	fmt.Println("Informasi tentang Book 1")
	fmt.Println("Judul : ", book1.Title)
	fmt.Println("Penulis : ", book1.Author)

	// nomor 2
	var book2 Book
	book2.Title = "Algoritma dan Struktur Data"
	book2.Author = "Jane Smith"

	fmt.Println("Informasi tentang Book 2")
	fmt.Println("Judul : ", book2.Title)
	fmt.Println("Penulis : ", book2.Author)

	//nomor 3
	book3 := Book{"Machine Learning for Beginners", "David Johnson"}

	fmt.Println("Informasi tentang Book 3")
	fmt.Println("Judul : ", book3.Title)
	fmt.Println("Penulis : ", book3.Author)

}

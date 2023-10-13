package main

import "fmt"

func main() {
	//map kosong
	contacts := make(map[string]string)
	fmt.Println(contacts)

	//menambahkan dua kontak
	contacts["Alice"] = "1234567890"
	contacts["Bob"] = "9876543210"

	//menampilkan semua kontak
	fmt.Println("Semua Kontak", contacts)
	fmt.Println("Nomor telepon Alice", contacts["Alice"])

	//menambahkan 1 kontak
	contacts["Charlie"] = "5555555555"
	fmt.Println("Setelah menambah kontak Charlie", contacts)

	//menghapus kontak
	delete(contacts, "Alice")
	fmt.Println("Setelah hapus kontak alice", contacts)

	//menampilkan semua kontak
	fmt.Println("Daftar Kontak")
	for key, value := range contacts {
		fmt.Printf("%s:%s \n", key, value)
	}
}

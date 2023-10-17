package main

import "fmt"

//membuat struct
type Siswa struct {
	OrangTua OrangTua
	Nama     string
	Umur     int
	Kelas    string
}

type OrangTua struct {
	Nama string
	Umur int
}

func main() {

	// membuat variabel siswa 1
	parent1 := OrangTua{
		Nama: "Budi",
		Umur: 40,
	}

	siswa1 := Siswa{
		OrangTua: parent1,
		Nama:     "Ali",
		Umur:     15,
		Kelas:    "9A",
	}

	// membuat variabel siswa 2
	parent2 := OrangTua{
		Nama: "Citra",
		Umur: 39,
	}

	siswa2 := Siswa{
		OrangTua: parent2,
		Nama:     "David",
		Umur:     14,
		Kelas:    "8B",
	}

	fmt.Println("Informasi Siswa 1:")
	fmt.Println("Nama: ", siswa1.Nama, ", Umur: ", siswa1.Umur, ", Kelas: ", siswa1.Kelas)
	fmt.Println("Orang Tua:", siswa1.OrangTua.Nama, ", Umur: ", siswa1.OrangTua.Umur)

	fmt.Println("Informasi Siswa 2:")
	fmt.Println("Nama: ", siswa2.Nama, ", Umur: ", siswa2.Umur, ", Kelas: ", siswa2.Kelas)
	fmt.Println("Orang Tua:", siswa2.OrangTua.Nama, ", Umur: ", siswa2.OrangTua.Umur)

	// membuat variable siswa 3
	parent3 := OrangTua{"Eva", 35}
	siswa3 := Siswa{parent3, "Fina", 12, "10C"}

	fmt.Println("Informasi Siswa 3:")
	fmt.Println("Nama: ", siswa3.Nama, ", Umur: ", siswa3.Umur, ", Kelas: ", siswa3.Kelas)
	fmt.Println("Orang Tua:", siswa3.OrangTua.Nama, ", Umur: ", siswa3.OrangTua.Umur)

}

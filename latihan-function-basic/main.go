package main

import "fmt"

type Car struct {
	Name  string
	Color string
}

func carInfo(car map[string]string) {
	name := car.Name
	color := car.color
	text := fmt.Println("Mobil %s berwarna %s", name, color)
	return text

}
func main() {
	var car = map[string]string
	car["name"] = "BWM"
	car["color"] = "Black"

	// buat 2 buah fungsi :
	// 1 => fungsi yang mengembalikan sebuah string
	// pada fungsi ini terjadi pengolahan kata sehingga menghasilkan kata : Mobil BMW berwarna Black

	// 2 => fungsi yang menampilkan hasil dari kembalian string
	// fungsi ini hanya bertugas untuk menampilkan kata

	// alur
	// simpan hasil dari return function kedalam sebuah variable message
	// tampilkan hasil dari variable message

	// output => Mobil BMW berwarna Black

}

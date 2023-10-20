package main

import (
	"fmt"
)

type Car struct {
	data map[string]string
}

// cara pertama
func carInfo(car map[string]string) string {
	name := car["name"]
	color := car["color"]
	text := fmt.Sprintf("Mobil %s berwarna %s\n", name, color)
	return text
}

func tampilCar(message string) {
	fmt.Println(message)
}

// cara kedua
func carInfo1() string {
	car := make(map[string]string)
	car["name"] = "BMW"
	car["color"] = "Black"
	text := fmt.Sprintf("Mobil %s Berwarna %s\n", car["name"], car["color"])
	return text
}

func main() {

	car := map[string]string{
		"name":  "BMW",
		"color": "Black",
	}

	//menampilkan hasil cara 1
	message := carInfo(car)
	fmt.Println(message)

	//menampilkan hasil cara 2
	message1 := carInfo1()
	fmt.Println(message1)

}

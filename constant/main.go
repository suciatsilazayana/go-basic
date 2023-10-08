package main

import "fmt"

// membuat constant diluar fungsi
const LENGTH int = 10

func main() {
	// didalam fungsi
	// constanta tidak dapat dubah jika sudah
	// dideklarasikan diawal
	const WIDTH = 5
	const PI = 3.14

	fmt.Println(LENGTH)
	fmt.Println(WIDTH)

	// deklarasi multiple
	const (
		x int = 10
		y     = 3.14
		z     = "Hello World"
	)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}

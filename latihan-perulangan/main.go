package main

import "fmt"

func main() {

	//soal 1
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}

	//soal 2
	for i := 5; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}

	//soal 3
	for i := 1; i <= 50; i++ {
		habis_dibagi_3 := i % 3
		habis_dibagi_5 := i % 5

		if habis_dibagi_3 == 0 && habis_dibagi_5 == 0 {
			fmt.Print("Noobee, ")
		} else if habis_dibagi_3 == 0 {
			fmt.Print("Noo, ")
		} else if habis_dibagi_5 == 0 {
			fmt.Print("Bee, ")
		} else {
			fmt.Print(i, ", ")
		}
	}

}

package main

import (
	"fmt"
	"log"
)

func main() {
	sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8}
	find := 8

	found := Binary(find, sortedArray...)
	fmt.Println(found)
}

func Binary(find int, sortedNums ...int) (found bool) {
	if len(sortedNums) <= 1 {
		if sortedNums[0] == find {
			return true
		}
		return false
	}

	// infinite loop
	// tujuannya untuk melakukan pencarian terus menerus
	for {
		// untuk mendapatkan index tengah nya
		middleIndex := len(sortedNums) / 2

		log.Println("middle index", middleIndex, "sorted nums", sortedNums, "target data", find)

		// cek apakah data di index ini == find
		if sortedNums[middleIndex] == find {
			log.Println("done with middle index", middleIndex, "sorted nums", sortedNums, "target data", find)
			return true
		} else if find < sortedNums[middleIndex] {
			// fungsi ini untuk mengubah sortedNums
			// yang mana artinya mengambil value
			// dari index 0 sampai index ke `middleIndex`
			sortedNums = sortedNums[:middleIndex]
		} else if find > sortedNums[middleIndex] {
			// fungsi ini untuk mengubah sortedNums
			// dari index `middleIndex` ke akhir index
			sortedNums = sortedNums[middleIndex:]
		}

		// jika data sudah tidak bisa di bagi lagi
		// maka keluar dari looping
		if middleIndex == 0 {
			break
		}
	}

	return false
}

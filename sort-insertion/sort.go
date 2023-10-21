package main

import "log"

func main() {
	sorted := InsertionSort(6, 5, 3, 1, 8, 7, 2, 4)
	log.Println(sorted)
}

// fungsi sebelumnya
// func BubbleSort(nums ...int) (sorted []int) {}

/**
fungsi ini sama dengan BubbleSort,
cuma beda di algorritma yang digunakan
*/
func InsertionSort(nums ...int) (sorted []int) {

	// index untuk titik awalnya adalah 1.
	// karena akan dibandingkan ke kiri, yaitu ke index 0
	for i := 1; i < len(nums); i++ {
		// index nya dimulai sama dengan index point, yaitu i
		// proses nya akan selalu berkurang ke kiri
		for k := i; k > 0; k-- {

			//log
			log.Println(nums[k-1], ">", nums[k])

			// proses untuk pemindahanya ada disini
			// logic nya adalah cek nilai 1 index sebelum index saat ini
			// lalu lakukan perbandingan
			if nums[k-1] > nums[k] {
				//log
				log.Println("change", nums[k-1], "in position", k-1, "to", nums[k], "in position", k)

				// proses swap value
				temp := nums[k]
				nums[k] = nums[k-1]
				nums[k-1] = temp

				log.Println(nums)
			}
		}
		log.Println("hasil sort batch", i+1, ":", nums)
	}

	sorted = nums

	return
}

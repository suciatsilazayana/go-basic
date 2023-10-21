package main

import "log"

func main() {
	sorted := BubbleSort(3, 1, 4, 2)
	log.Println("Sorted array", sorted)
}

// fungsi untuk melakukan sorting menggunakan bubble sort
// parameternya adalah variadic, untuk memudahkan
// proses pengiriman value nantinya.
// contoh : BubbleSort(3,1,4,2)
func BubbleSort(nums ...int) (sorted []int) {

	/*
	   kode untuk menentukan titik point nya
	   jadi index i, akan bertugas sebagai point yang dibandingkan.
	   lalu < len(nums)-1, karena kita hanya akan meletakkan titik point
	   di 1 index sebelum terakhir. Karena index terakhir
	   akan diisi oleh pembandingnya
	*/
	for i := 0; i < len(nums)-1; i++ {

		/*
		   kode untuk menentukan pembandingnya
		   jadi index j, akan bertugas untuk melakukan perbandingan
		   ke titik point (index i). Index j akan
		   selalu lebih dulu (i+1) dibanding index i.
		*/
		for j := i + 1; j < len(nums); j++ {

			// log operasi yang akan digunakan
			log.Println(nums[i], ">", nums[j])

			// jika kondisi benar, maka akan kita swap posisi nya
			if nums[i] > nums[j] {
				// log perubahan posisi nya
				log.Println("change", nums[i], "in position", i, "to", nums[j], "in position", j)

				// simpan value nums[i] kedalam variable
				// temporary / sementara
				temp := nums[i]

				// ubah value index i ke index j
				nums[i] = nums[j]

				// lalu ubah value index j ke temp
				// (isi nya adalah value dari index i sebelum berubah)
				nums[j] = temp

				// log kondisi sekarang
				log.Println(nums)
			}
		}

		log.Println("hasil sort batch", i+1, ":", nums)

	}
	return
}

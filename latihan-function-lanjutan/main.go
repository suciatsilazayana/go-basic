package main

import "fmt"

func sum2(nums ...float32) (total float32) {
	fmt.Println(nums)

	for _, num := range nums {
		total += num // total = total + num
	}
	return total
}

func ratarata(nilaiSiswa ...float32) float32 {
	total := float32(0)

	for _, nilai := range nilaiSiswa {
		total += nilai
	}

	return total / float32(len(nilaiSiswa))
}
func main() {
	nilaiSiswa := []float32{85.5, 92.0, 78.5, 90.0, 88.5}

	fmt.Println("Jumlah nilai siswa dalam kelas:", len(nilaiSiswa))
	fmt.Println("Nilai rata-rata siswa dalam kelas:", ratarata(nilaiSiswa...))

}

package main

import (
	"fmt"
)

type Examination struct {
	Name  string
	Score int
}

func getMedian(data []Examination) int {
	var median int

	n := len(data)
	if n%2 == 0 {
		median = (data[n/2-1].Score + data[n/2].Score) / 2
		return median
	} else {
		median = data[n/2].Score
		return median
	}
}

func getMax(data []Examination) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0].Score

	for _, item := range data {
		if item.Score > max {
			max = item.Score
		}
	}
	return max
}

func getMin(data []Examination) int {
	if len(data) == 0 {
		return 0
	}

	min := data[0].Score

	for _, item := range data {
		if item.Score < min {
			min = item.Score
		}
	}

	return min
}

func getAverage(data []Examination) float64 {
	total := 0
	for _, exam := range data {
		total += exam.Score
	}
	return float64(total) / float64(len(data))
}

// func Binary(find int, Examination ...int) (found bool) {
// 	if len(Examination) <= 1 {
// 		if Examination[0] == find {
// 			return true
// 		}
// 		return false
// 	}

// 	for {

// 		middleIndex := len(Examination) / 2

// 		log.Println("middle index", middleIndex, "sorted nums", Examination, "target data", find)

// 		if Examination[middleIndex] == find {
// 			log.Println("done with middle index", middleIndex, "sorted nums", Examination, "target data", find)
// 			return true
// 		} else if find < Examination[middleIndex] {

// 			Examination = Examination[:middleIndex]
// 		} else if find > Examination[middleIndex] {

// 			Examination = Examination[middleIndex:]
// 		}
// 		if middleIndex == 0 {
// 			break
// 		}
// 	}

// 	return false
// }

func main() {
	examResult := []Examination{
		{Name: "Budi",
			Score: 90},
		{Name: "Andi",
			Score: 85},
		{Name: "Nayla",
			Score: 87},
		{Name: "Danu",
			Score: 80},
		{Name: "Rahmat",
			Score: 75},
		{Name: "Erika",
			Score: 83},
		{Name: "Siska",
			Score: 87},
		{Name: "Mita",
			Score: 94},
		{Name: "Shinta",
			Score: 82},
		{Name: "Jojo",
			Score: 83},
	}

	median := getMedian(examResult)
	max := getMax(examResult)
	min := getMin(examResult)
	average := getAverage(examResult)
	fmt.Println("Nilai median:", median)
	fmt.Println("Nilai paling tinggi:", max)
	fmt.Println("Nilai paling rendah", min)
	fmt.Println("Nilai rata rata", average)
	// find := 80

	// found := Binary(find, examResult...)
	// fmt.Println(found)

}

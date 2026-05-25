package main

import "fmt"

func main() {
	var n int
	fmt.Print("Jumlah data: ")
	fmt.Scan(&n)

	var arr [100]float64
	sum := 0.0

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		sum += arr[i]
	}

	min := arr[0]
	max := arr[0]

	for i := 1; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	rata := sum / float64(n)

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
	fmt.Println("Rata-rata:", rata)
}

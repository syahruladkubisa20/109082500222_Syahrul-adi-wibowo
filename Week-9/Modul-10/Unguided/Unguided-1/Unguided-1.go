package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(&n)

	var arr [100]float64

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
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

	fmt.Println("Minimum:", min)
	fmt.Println("Maximum:", max)
}

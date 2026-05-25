package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var arr [100]float64

	for i := 0; i < x; i++ {
		fmt.Scan(&arr[i])
	}

	total := 0.0
	count := 0

	for i := 0; i < x; i++ {
		total += arr[i]
		count++
		if count == y {
			fmt.Println("Total wadah:", total)
			total = 0
			count = 0
		}
	}
}

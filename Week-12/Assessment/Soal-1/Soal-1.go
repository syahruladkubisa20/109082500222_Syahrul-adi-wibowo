package main

import "fmt"

func SelectionSort(a *[1000]int, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			a[i], a[minIdx] = a[minIdx], a[i]
		}
	}
}

func median(a [1000]int, n int) float64 {
	if n == 0 {
		return 0
	}
	if n%2 == 1 {
		return float64(a[n/2])
	}
	mid1 := a[(n/2)-1]
	mid2 := a[n/2]
	return float64(mid1+mid2) / 2.0
}

func main() {
	var data [1000]int
	n := 0
	for {
		var x int
		if _, err := fmt.Scan(&x); err != nil {
			break
		}
		if x == -5313541 {
			break
		}
		if x == 0 {
			if n == 0 {
				fmt.Println("Data kosong")
				continue
			}
			SelectionSort(&data, n)
			fmt.Print("Data terurut: ")
			for i := 0; i < n; i++ {
				if i > 0 {
					fmt.Print(" ")
				}
				fmt.Print(data[i])
			}
			fmt.Println()
			m := median(data, n)
			if m == float64(int64(m)) {
				fmt.Printf("Median: %.0f\n", m)
			} else {
				fmt.Printf("Median: %g\n", m)
			}
			continue
		}
		if n < len(data) {
			data[n] = x
			n++
		}
	}
}

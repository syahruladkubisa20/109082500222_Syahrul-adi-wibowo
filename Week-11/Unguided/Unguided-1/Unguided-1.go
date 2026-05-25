package main

import "fmt"

func main() {
	var suara [21]int
	var masuk, sah int
	for {
		var x int
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		masuk++
		if x >= 1 && x <= 20 {
			suara[x]++
			sah++
		}
	}
	fmt.Println("Jumlah suara masuk:", masuk)
	fmt.Println("Jumlah suara sah:", sah)
	fmt.Println("Daftar calon yang mendapat suara:")
	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("Calon %d: %d suara\n", i, suara[i])
		}
	}
}

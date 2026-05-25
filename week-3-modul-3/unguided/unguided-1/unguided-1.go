package main

import "fmt"

func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func main() {
	var angka int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	fmt.Println("Hasil faktorial:", faktorial(angka))
}

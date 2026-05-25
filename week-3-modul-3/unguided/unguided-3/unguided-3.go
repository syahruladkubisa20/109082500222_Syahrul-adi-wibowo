package main

import "fmt"

func tampilDeret(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}

func main() {
	var angka int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	tampilDeret(angka)
}

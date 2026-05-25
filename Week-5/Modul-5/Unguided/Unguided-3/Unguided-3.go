package main

import "fmt"

func tampilGanjil(n int) {
	if n < 1 {
		return
	}
	tampilGanjil(n - 1)

	if n%2 != 0 {
		fmt.Print(n, " ")
	}
}

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	fmt.Print("Bilangan ganjil: ")
	tampilGanjil(n)
}

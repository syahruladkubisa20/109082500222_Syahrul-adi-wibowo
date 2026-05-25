package main

import "fmt"

func tampilFaktor(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	tampilFaktor(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	fmt.Print("Faktor: ")
	tampilFaktor(n, 1)
}

package main

import "fmt"

func cekGenap(angka int) bool {
	if angka%2 == 0 {
		return true
	}

	return false
}

func main() {
	angka := 8

	hasilGenap := cekGenap(angka)
	if hasilGenap {
		fmt.Printf("%d adalah angka genap\n", angka)
	} else {
		fmt.Printf("%d adalah angka ganjil\n", angka)
	}
}

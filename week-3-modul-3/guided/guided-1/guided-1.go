package main

import "fmt"

func HitungLuasPersegiPanjang(panjang int, lebar int) int {
	luas := panjang * lebar
	return luas
}

func main() {
	panjang := 10
	lebar := 5

	luas := HitungLuasPersegiPanjang(panjang, lebar)
	fmt.Printf("Luas persegi panjang dengan panjang %d dan lebar %d adalah %d\n", panjang, lebar, luas)
}

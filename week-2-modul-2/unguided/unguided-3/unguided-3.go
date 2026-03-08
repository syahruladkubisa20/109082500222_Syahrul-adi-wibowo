package main

import "fmt"

func hitungBiaya(gram int) {

	kg := gram / 1000
	sisa := gram % 1000

	biayaKg := kg * 10000
	biayaSisa := 0

	if kg > 10 {
		biayaSisa = 0
	} else {
		if sisa > 500 {
			biayaSisa = sisa * 5
		} else {
			biayaSisa = sisa * 15
		}
	}

	total := biayaKg + biayaSisa

	fmt.Println("Berat:", gram, "gram")
	fmt.Println("Detail:", kg, "kg +", sisa, "gram")
	fmt.Println("Biaya kg:", biayaKg)
	fmt.Println("Biaya sisa:", biayaSisa)
	fmt.Println("Total biaya:", total)
	fmt.Println()
}

func main() {

	hitungBiaya(8500)
	hitungBiaya(9250)
	hitungBiaya(11750)

}

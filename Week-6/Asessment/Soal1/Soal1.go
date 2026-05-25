package main

import "fmt"

const pi = 3.14

func main() {
	var r, tKiri, tKanan, pKiri, pKanan float64

	fmt.Print("jari-jari alas tabung : ")
	fmt.Scan(&r)

	fmt.Print("tinggi zat cair tabung kiri : ")
	fmt.Scan(&tKiri)
	fmt.Print("massa jenis zat cair tabung kiri : ")
	fmt.Scan(&pKiri)

	fmt.Print("tinggi zat cair tabung kanan : ")
	fmt.Scan(&tKanan)
	fmt.Print("massa jenis zat cair tabung kanan : ")
	fmt.Scan(&pKanan)

	massaKiri := pi * r * r * tKiri * pKiri
	massaKanan := pi * r * r * tKanan * pKanan

	if massaKiri == massaKanan {
		fmt.Println("BALANCE")
	} else {
		selisih := massaKiri - massaKanan
		if selisih < 0 {
			selisih = -selisih
		}
		fmt.Printf("Selisih massa zat cair kiri dan kanan: %.0f\n", selisih)
	}
}

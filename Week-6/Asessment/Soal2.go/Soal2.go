package main

import "fmt"

func main() {
	var mhs, hari int
	var tujuan string

	fmt.Print("Jumlah mahasiswa: ")
	fmt.Scan(&mhs)
	fmt.Print("Lama hari study tour: ")
	fmt.Scan(&hari)
	fmt.Print("Tujuan (domestik/mancanegara): ")
	fmt.Scan(&tujuan)

	if tujuan == "domestik" && hari > 3 {
		hari = 3
	}
	if tujuan == "mancanegara" && hari > 8 {
		hari = 8
	}

	biayaHarian := 35000*2 + 250000 + 300000
	if tujuan == "mancanegara" {
		biayaHarian = int(float64(biayaHarian) * 1.5)
	}

	total := mhs * hari * biayaHarian
	fmt.Println("Total biaya yang harus dikeluarkan Tel-U : Rp.", total)
}

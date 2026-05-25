package main

import "fmt"

type NamaProv [11]string
type PopProv [11]int
type TumbuhProv [11]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	idx := 1
	for i := 2; i <= 10; i++ {
		if tumbuh[i] > tumbuh[idx] {
			idx = i
		}
	}
	return idx
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 1; i <= 10; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("\n=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")
	for i := 1; i <= 10; i++ {
		if tumbuh[i] > 0.02 {
			hasil := float64(pop[i]) * (tumbuh[i] + 1)
			fmt.Printf("%s %.0f\n", prov[i], hasil)
		}
	}
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv
	var cari string

	fmt.Println("=== Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ===")
	InputData(&prov, &pop, &tumbuh)
	
	fmt.Scan(&cari)

	idxCepat := ProvinsiTercepat(tumbuh)
	fmt.Printf("\nProvinsi dengan angka pertumbuhan tercepat : %s\n", prov[idxCepat])

	fmt.Printf("\nData provinsi yang dicari : %s\n", cari)

	Prediksi(prov, pop, tumbuh)
}
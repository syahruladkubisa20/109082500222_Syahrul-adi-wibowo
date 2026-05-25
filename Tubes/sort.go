package main

import "fmt"

// Selection Sort berdasarkan tanggal jatuh tempo
func selectionSortTanggal() {
	n := len(daftarTagihan)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if daftarTagihan[j].JatuhTempo < daftarTagihan[minIdx].JatuhTempo {
				minIdx = j
			}
		}
		if minIdx != i {
			daftarTagihan[i], daftarTagihan[minIdx] = daftarTagihan[minIdx], daftarTagihan[i]
		}
	}
	fmt.Println("Data berhasil diurutkan berdasarkan tanggal jatuh tempo (Selection Sort).")
}

// Insertion Sort berdasarkan nominal
func insertionSortNominal() {
	n := len(daftarTagihan)
	for i := 1; i < n; i++ {
		key := daftarTagihan[i]
		j := i - 1
		for j >= 0 && daftarTagihan[j].Nominal > key.Nominal {
			daftarTagihan[j+1] = daftarTagihan[j]
			j--
		}
		daftarTagihan[j+1] = key
	}
	fmt.Println("Data berhasil diurutkan berdasarkan nominal (Insertion Sort).")
}

func urutkanTagihan() {
	fmt.Println("\n--- Urutkan Tagihan ---")
	fmt.Println("1. Urutkan berdasarkan Tanggal Jatuh Tempo (Selection Sort)")
	fmt.Println("2. Urutkan berdasarkan Nominal (Insertion Sort)")
	fmt.Print("Pilih: ")
	var pilihan int
	fmt.Scan(&pilihan)
	scanner.Scan()

	switch pilihan {
	case 1:
		selectionSortTanggal()
		cetakTabel()
	case 2:
		insertionSortNominal()
		cetakTabel()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

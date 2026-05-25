package main

import "fmt"

func statistikTagihan() {
	if len(daftarTagihan) == 0 {
		fmt.Println("Belum ada data tagihan.")
		return
	}

	var totalSemua float64
	var totalLunas float64
	var jumlahLunas int

	for _, t := range daftarTagihan {
		totalSemua += t.Nominal
		if t.Status == "Lunas" {
			totalLunas += t.Nominal
			jumlahLunas++
		}
	}

	persenLunas := 0.0
	if len(daftarTagihan) > 0 {
		persenLunas = float64(jumlahLunas) / float64(len(daftarTagihan)) * 100
	}

	fmt.Println("\n========================================")
	fmt.Println("         STATISTIK TAGIHAN             ")
	fmt.Println("========================================")
	fmt.Printf("Total semua tagihan   : Rp %.0f\n", totalSemua)
	fmt.Printf("Total tagihan lunas   : Rp %.0f\n", totalLunas)
	fmt.Printf("Total tagihan belum   : Rp %.0f\n", totalSemua-totalLunas)
	fmt.Printf("Jumlah tagihan        : %d\n", len(daftarTagihan))
	fmt.Printf("Jumlah lunas          : %d\n", jumlahLunas)
	fmt.Printf("Jumlah belum lunas    : %d\n", len(daftarTagihan)-jumlahLunas)
	fmt.Printf("Persentase lunas      : %.2f%%\n", persenLunas)
	fmt.Println("========================================")
}

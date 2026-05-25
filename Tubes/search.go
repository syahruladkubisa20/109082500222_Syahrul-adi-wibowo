package main

import (
	"fmt"
	"strings"
)

// Sequential Search berdasarkan nama tagihan
func sequentialSearch(keyword string) {
	fmt.Println("\n--- Hasil Pencarian (Sequential Search) ---")
	keyword = strings.ToLower(keyword)
	ketemu := false
	for i, t := range daftarTagihan {
		if strings.Contains(strings.ToLower(t.Nama), keyword) {
			fmt.Printf("%-4d %-20s %-12s %-12.0f %-12s %-8s\n", i+1, t.Nama, t.Kategori, t.Nominal, t.JatuhTempo, t.Status)
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("Tagihan tidak ditemukan.")
	}
}

// Binary Search berdasarkan kategori (data harus sudah diurutkan dulu)
func binarySearchKategori(keyword string) {
	fmt.Println("\n--- Hasil Pencarian (Binary Search by Kategori) ---")

	// salin dulu biar data asli tidak berubah
	temp := make([]Tagihan, len(daftarTagihan))
	copy(temp, daftarTagihan)

	// urutkan dulu berdasarkan kategori (bubble sort sederhana)
	for i := 0; i < len(temp)-1; i++ {
		for j := 0; j < len(temp)-i-1; j++ {
			if strings.ToLower(temp[j].Kategori) > strings.ToLower(temp[j+1].Kategori) {
				temp[j], temp[j+1] = temp[j+1], temp[j]
			}
		}
	}

	keyword = strings.ToLower(keyword)
	low := 0
	high := len(temp) - 1
	ketemu := false

	for low <= high {
		mid := (low + high) / 2
		kat := strings.ToLower(temp[mid].Kategori)
		if kat == keyword {
			// cari semua yang sama kategorinya
			// cek ke kiri
			start := mid
			for start > 0 && strings.ToLower(temp[start-1].Kategori) == keyword {
				start--
			}
			// cetak semua yang cocok
			for k := start; k < len(temp) && strings.ToLower(temp[k].Kategori) == keyword; k++ {
				fmt.Printf("%-4d %-20s %-12s %-12.0f %-12s %-8s\n", k+1, temp[k].Nama, temp[k].Kategori, temp[k].Nominal, temp[k].JatuhTempo, temp[k].Status)
				ketemu = true
			}
			break
		} else if kat < keyword {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !ketemu {
		fmt.Println("Kategori tidak ditemukan.")
	}
}

func cariTagihan() {
	fmt.Println("\n--- Cari Tagihan ---")
	fmt.Println("1. Cari berdasarkan Nama (Sequential Search)")
	fmt.Println("2. Cari berdasarkan Kategori (Binary Search)")
	fmt.Print("Pilih: ")
	var pilihan int
	fmt.Scan(&pilihan)
	scanner.Scan()

	switch pilihan {
	case 1:
		keyword := bacaInput("Masukkan nama tagihan: ")
		sequentialSearch(keyword)
	case 2:
		keyword := bacaInput("Masukkan kategori: ")
		binarySearchKategori(keyword)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

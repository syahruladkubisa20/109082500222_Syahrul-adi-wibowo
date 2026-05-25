package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func bacaInput(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func tampilkanMenu() {
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("              MENU UTAMA               ")
	fmt.Println("========================================")
	fmt.Println("1. Tambah Tagihan")
	fmt.Println("2. Lihat Daftar Tagihan")
	fmt.Println("3. Edit Tagihan")
	fmt.Println("4. Hapus Tagihan")
	fmt.Println("5. Cari Tagihan")
	fmt.Println("6. Urutkan Tagihan")
	fmt.Println("7. Statistik Tagihan")
	fmt.Println("0. Exit")
	fmt.Println("========================================")
}

func cetakTabel() {
	fmt.Println()
	fmt.Printf("%-4s %-20s %-12s %-12s %-12s %-8s\n", "No", "Nama Tagihan", "Kategori", "Nominal", "Jatuh Tempo", "Status")
	fmt.Println("------------------------------------------------------------------------")
	for i, t := range daftarTagihan {
		fmt.Printf("%-4d %-20s %-12s %-12.0f %-12s %-8s\n", i+1, t.Nama, t.Kategori, t.Nominal, t.JatuhTempo, t.Status)
	}
	fmt.Println("------------------------------------------------------------------------")
}

func tambahTagihan() {
	fmt.Println("\n--- Tambah Tagihan Baru ---")
	nama := bacaInput("Nama tagihan: ")
	if nama == "" {
		fmt.Println("Nama tidak boleh kosong!")
		return
	}
	kategori := bacaInput("Kategori (Utilitas/Hiburan/Kesehatan/Elektronik/Lainnya): ")
	if kategori == "" {
		fmt.Println("Kategori tidak boleh kosong!")
		return
	}

	var nominal float64
	fmt.Print("Nominal (Rp): ")
	fmt.Scan(&nominal)
	// consume newline
	scanner.Scan()

	if nominal <= 0 {
		fmt.Println("Nominal harus lebih dari 0!")
		return
	}

	jatuhTempo := bacaInput("Jatuh tempo (YYYY-MM-DD): ")
	if jatuhTempo == "" {
		fmt.Println("Jatuh tempo tidak boleh kosong!")
		return
	}

	status := bacaInput("Status (Lunas/Belum): ")
	if status != "Lunas" && status != "Belum" {
		fmt.Println("Status harus 'Lunas' atau 'Belum'!")
		return
	}

	tagihan := Tagihan{
		Nama:       nama,
		Kategori:   kategori,
		Nominal:    nominal,
		JatuhTempo: jatuhTempo,
		Status:     status,
	}
	daftarTagihan = append(daftarTagihan, tagihan)
	fmt.Println("Tagihan berhasil ditambahkan!")
}

func lihatTagihan() {
	if len(daftarTagihan) == 0 {
		fmt.Println("Belum ada data tagihan.")
		return
	}
	fmt.Println("\n--- Daftar Tagihan ---")
	cetakTabel()
}

func editTagihan() {
	if len(daftarTagihan) == 0 {
		fmt.Println("Belum ada data tagihan.")
		return
	}
	cetakTabel()
	var no int
	fmt.Print("Masukkan nomor tagihan yang ingin diedit: ")
	fmt.Scan(&no)
	scanner.Scan()

	if no < 1 || no > len(daftarTagihan) {
		fmt.Println("Nomor tidak valid!")
		return
	}

	idx := no - 1
	fmt.Println("Kosongkan input jika tidak ingin mengubah.")

	nama := bacaInput("Nama baru (" + daftarTagihan[idx].Nama + "): ")
	if nama != "" {
		daftarTagihan[idx].Nama = nama
	}

	kategori := bacaInput("Kategori baru (" + daftarTagihan[idx].Kategori + "): ")
	if kategori != "" {
		daftarTagihan[idx].Kategori = kategori
	}

	var nominal float64
	fmt.Printf("Nominal baru (%.0f), masukkan 0 untuk skip: ", daftarTagihan[idx].Nominal)
	fmt.Scan(&nominal)
	scanner.Scan()
	if nominal > 0 {
		daftarTagihan[idx].Nominal = nominal
	}

	jatuhTempo := bacaInput("Jatuh tempo baru (" + daftarTagihan[idx].JatuhTempo + "): ")
	if jatuhTempo != "" {
		daftarTagihan[idx].JatuhTempo = jatuhTempo
	}

	status := bacaInput("Status baru (" + daftarTagihan[idx].Status + ") (Lunas/Belum): ")
	if status == "Lunas" || status == "Belum" {
		daftarTagihan[idx].Status = status
	}

	fmt.Println("Tagihan berhasil diupdate!")
}

func hapusTagihan() {
	if len(daftarTagihan) == 0 {
		fmt.Println("Belum ada data tagihan.")
		return
	}
	cetakTabel()
	var no int
	fmt.Print("Masukkan nomor tagihan yang ingin dihapus: ")
	fmt.Scan(&no)
	scanner.Scan()

	if no < 1 || no > len(daftarTagihan) {
		fmt.Println("Nomor tidak valid!")
		return
	}

	nama := daftarTagihan[no-1].Nama
	daftarTagihan = append(daftarTagihan[:no-1], daftarTagihan[no:]...)
	fmt.Printf("Tagihan '%s' berhasil dihapus!\n", nama)
}

package main

import "fmt"

type angka int
type kata string

type Buku struct {
	judul         kata
	penulis       kata
	penerbit      kata
	tahunTerbit   angka
	jumlahHalaman angka
}

func main() {
	var b Buku

	fmt.Println("=== INPUT BIODATA BUKU ===")

	fmt.Print("Judul: ")
	fmt.Scan(&b.judul)

	fmt.Print("Penulis: ")
	fmt.Scan(&b.penulis)

	fmt.Print("Penerbit: ")
	fmt.Scan(&b.penerbit)

	fmt.Print("Tahun Terbit: ")
	fmt.Scan(&b.tahunTerbit)

	fmt.Print("Jumlah Halaman: ")
	fmt.Scan(&b.jumlahHalaman)

	fmt.Println("\n=== DATA BUKU ===")
	fmt.Println("Judul          :", b.judul)
	fmt.Println("Penulis        :", b.penulis)
	fmt.Println("Penerbit       :", b.penerbit)
	fmt.Println("Tahun Terbit   :", b.tahunTerbit)
	fmt.Println("Jumlah Halaman :", b.jumlahHalaman)
}

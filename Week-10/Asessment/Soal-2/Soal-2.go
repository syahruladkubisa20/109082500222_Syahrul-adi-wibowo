package main

import "fmt"

type mahasiswa struct {
	nim, nama string
	nilai     int
}

type arrayMahasiswa [52]mahasiswa

func cariNilaiPertama(T arrayMahasiswa, n int, target string) int {
	for i := 1; i <= n; i++ {
		if T[i].nim == target {
			return T[i].nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(T arrayMahasiswa, n int, target string) int {
	max := -1
	for i := 1; i <= n; i++ {
		if T[i].nim == target {
			if T[i].nilai > max {
				max = T[i].nilai
			}
		}
	}
	return max
}

func main() {
	var T arrayMahasiswa
	var n int
	var target string

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i)
		fmt.Scan(&T[i].nim, &T[i].nama, &T[i].nilai)
	}

	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&target)

	p1 := cariNilaiPertama(T, n, target)
	p2 := cariNilaiTerbesar(T, n, target)

	fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", target, p1)
	fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", target, p2)
}

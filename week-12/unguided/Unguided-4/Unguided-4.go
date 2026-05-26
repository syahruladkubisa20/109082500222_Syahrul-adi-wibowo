package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Buku struct {
	id       string
	judul    string
	penulis  string
	penerbit string
	eksemplar int
	tahun    int
	rating   int
}

const nMax = 7919
var pustaka [nMax]Buku
var nPustaka int

func DaftarkanBuku() {
	fmt.Print("")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	nPustaka = n
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		data := strings.SplitN(line, " ", 7)
		pustaka[i].id = data[0]
		pustaka[i].judul = data[1]
		pustaka[i].penulis = data[2]
		pustaka[i].penerbit = data[3]
		pustaka[i].eksemplar, _ = strconv.Atoi(data[4])
		pustaka[i].tahun, _ = strconv.Atoi(data[5])
		pustaka[i].rating, _ = strconv.Atoi(data[6])
	}
}

func CetakTerfavorit() {
	if nPustaka == 0 {
		fmt.Println("Tidak ada buku")
		return
	}
	maxIdx := 0
	for i := 1; i < nPustaka; i++ {
		if pustaka[i].rating > pustaka[maxIdx].rating {
			maxIdx = i
		}
	}
	b := pustaka[maxIdx]
	fmt.Printf("%s %s %s %d\n", b.judul, b.penulis, b.penerbit, b.tahun)
}

func UrutBuku() {
	for i := 1; i < nPustaka; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

func Cetak5Terbaru() {
	batas := 5
	if nPustaka < 5 {
		batas = nPustaka
	}
	for i := 0; i < batas; i++ {
		b := pustaka[i]
		fmt.Printf("%s %s %s %d\n", b.judul, b.penulis, b.penerbit, b.tahun)
	}
}

func CariBuku() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	r, _ := strconv.Atoi(scanner.Text())
	low, high := 0, nPustaka-1
	found := false
	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].rating == r {
			b := pustaka[mid]
			fmt.Printf("%s %s %s %d\n", b.judul, b.penulis, b.penerbit, b.tahun)
			found = true
			break
		} else if pustaka[mid].rating < r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if !found {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	DaftarkanBuku()
	CetakTerfavorit()
	UrutBuku()
	Cetak5Terbaru()
	CariBuku()
}

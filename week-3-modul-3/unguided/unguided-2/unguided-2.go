package main

import "fmt"

func kuadrat(x int) int {
	return x * x
}

func main() {
	var angka int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	fmt.Println("Hasil kuadrat:", kuadrat(angka))
}

package main

import "fmt"

func main() {
	var n int
	fmt.Print("Jumlah data: ")
	fmt.Scan(&n)

	var arr [100]int

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Isi array:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}

	fmt.Println("\nIndeks genap:")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}

	fmt.Println("\nIndeks ganjil:")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
}

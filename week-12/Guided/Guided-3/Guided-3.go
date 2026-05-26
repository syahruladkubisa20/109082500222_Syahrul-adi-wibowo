package main

import "fmt"

// Fungsi pass by value
func tambahValue(x int) {
	x = x + 10
	fmt.Printf("Nilai di dalam tambahValue: %d\n", x)
}

// Fungsi pass by reference (pointer)
func tambahReference(x *int) {
	*x = *x + 10
	fmt.Printf("Nilai di dalam tambahReference: %d\n", *x)
}

func main() {
	var a int = 5
	fmt.Printf("Nilai awal: %d\n", a)
	tambahValue(a)
	fmt.Printf("Setelah tambahValue: %d\n", a)
	tambahReference(&a)
	fmt.Printf("Setelah tambahReference: %d\n", a)
}

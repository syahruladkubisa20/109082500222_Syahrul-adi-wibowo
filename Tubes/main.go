package main

import "fmt"

func main() {
	fmt.Println("")
	fmt.Println("   SIMTAB - Sistem Manajemen Tagihan   ")
	fmt.Println("          Bulanan v1.0                 ")
	fmt.Println("")

	for {
		tampilkanMenu()
		pilihan := 0
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahTagihan()
		case 2:
			lihatTagihan()
		case 3:
			editTagihan()
		case 4:
			hapusTagihan()
		case 5:
			cariTagihan()
		case 6:
			urutkanTagihan()
		case 7:
			statistikTagihan()
		case 0:
			fmt.Println("Terima kasih sudah menggunakan SIMTAB!")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

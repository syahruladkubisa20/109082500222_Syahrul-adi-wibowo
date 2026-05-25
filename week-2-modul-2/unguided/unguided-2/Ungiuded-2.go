package main

import "fmt"

func main() {

	var g1, g2, g3, g4 string

	for i := 1; i <= 5; i++ {
		fmt.Println("Percobaan ke-", i)
		fmt.Scan(&g1, &g2, &g3, &g4)

		if g1 == "merah" && g2 == "kuning" && g3 == "hijau" && g4 == "ungu" {
			fmt.Println(true)
		} else {
			fmt.Println(false)
		}
	}
}

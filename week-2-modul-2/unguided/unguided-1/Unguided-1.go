package main

import "fmt"

func isKabisat(tahun int) bool {
	if tahun%400 == 0 {
		return true
	} else if tahun%4 == 0 && tahun%100 != 0 {
		return true
	}
	return false
}

func main() {
	tahun := []int{2000, 1974, 2106, 1876}

	for _, t := range tahun {
		fmt.Println(t, ":", isKabisat(t))
	}
}

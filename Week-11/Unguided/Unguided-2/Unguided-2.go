package main

import "fmt"

func main() {
	var suara [21]int
	for {
		var x int
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		if x >= 1 && x <= 20 {
			suara[x]++
		}
	}

	ketua, wakil := 0, 0
	for i := 1; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			wakil = ketua
			ketua = i
		} else if suara[i] > suara[wakil] && i != ketua {
			wakil = i
		}
	}
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil Ketua RT:", wakil)
}

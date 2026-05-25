package main

import "fmt"

type Partai struct {
	nama  string
	suara int
}

func sequentialSearch(list []Partai, n int, key string) int {
	for i := 0; i < n; i++ {
		if list[i].nama == key {
			return i
		}
	}
	return -1
}

func insertionSortDesc(list []Partai, n int) {
	for i := 1; i < n; i++ {
		key := list[i]
		j := i - 1
		for j >= 0 && list[j].suara < key.suara {
			list[j+1] = list[j]
			j--
		}
		list[j+1] = key
	}
}

func main() {
	var arr [100]Partai
	count := 0
	for {
		var nama string
		var suara int
		fmt.Scan(&nama)
		if nama == "-1" {
			break
		}
		fmt.Scan(&suara)
		idx := sequentialSearch(arr[:], count, nama)
		if idx != -1 {
			arr[idx].suara += suara
		} else {
			if count < len(arr) {
				arr[count] = Partai{nama: nama, suara: suara}
				count++
			}
		}
	}
	if count == 0 {
		return
	}
	list := arr[:count]
	insertionSortDesc(list, count)
	for i := 0; i < count; i++ {
		fmt.Printf("%s(%d)\n", list[i].nama, list[i].suara)
	}
}

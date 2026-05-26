
package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	var arr []int
	for {
		var x int
		fmt.Scan(&x)
		if x < 0 {
			break
		}
		arr = append(arr, x)
	}
	insertionSort(arr)
	for i, v := range arr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
	if len(arr) < 2 {
		fmt.Println("Data berjarak tidak tetap")
		return
	}
	jarak := arr[1] - arr[0]
	berjarak := true
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != jarak {
			berjarak = false
			break
		}
	}
	if berjarak {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}

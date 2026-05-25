package main

import "fmt"

func main() {
	var kata string
	fmt.Scan(&kata)

	var balik string

	for i := len(kata) - 1; i >= 0; i-- {
		balik += string(kata[i])
	}

	if kata == balik {
		fmt.Println("Palindrom")
	} else {
		fmt.Println("Bukan Palindrom")
	}
}

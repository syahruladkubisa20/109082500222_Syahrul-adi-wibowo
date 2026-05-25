package main

import "fmt"

func main() {
	var A, B string
	fmt.Scan(&A)
	fmt.Scan(&B)

	var a, b int

	for {
		fmt.Scan(&a, &b)

		if a < 0 || b < 0 {
			break
		}

		if a > b {
			fmt.Println(A)
		} else if b > a {
			fmt.Println(B)
		} else {
			fmt.Println("Draw")
		}
	}
}

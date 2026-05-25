package main

import "fmt"

type suhu float64

func CelciusToReamur(c suhu) suhu {
	return c * 4 / 5
}

func CelciusToFahrenheit(c suhu) suhu {
	return (c * 9 / 5) + 32
}

func CelciusToKelvin(c suhu) suhu {
	return c + 273.15
}

func main() {
	var input float64

	fmt.Println("=== KONVERTER SUHU ===")
	fmt.Print("Masukkan suhu dalam Celcius: ")
	fmt.Scan(&input)

	c := suhu(input)

	fmt.Printf("\n%.2f Celcius = %.2f Reamur\n", input, CelciusToReamur(c))
	fmt.Printf("%.2f Celcius = %.2f Fahrenheit\n", input, CelciusToFahrenheit(c))
	fmt.Printf("%.2f Celcius = %.2f Kelvin\n", input, CelciusToKelvin(c))
}

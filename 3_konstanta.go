package main

import "fmt"

func main()  {
	// Konstanta
	// Data seperti PI = 22,7, kecepatan cahaya
	const firstName string = "Ihsan"
	const lastName = "Iqbal"
	fmt.Print("Halo ", firstName, "!\n")

	// Multi-konstanta
	const (
		square	= "Kotak"
		isToday bool = true
		numeric = 8
		float = 1.5
		a = "nilai sama dengan b"
		b
	)
	const satu, dua = 1, 2
	fmt.Println(b)
}

package main

import "fmt"

func main()  {
	// Perulangan for
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// For dengan argumen hanya kondisi
	var i = 0

	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}
	fmt.Println()

	// For tanpa argumen
	var x = 0

	for {
		fmt.Println("Angka", x)

		x++
		if x == 5 {
			break
		}
	}
	fmt.Println()

	// Break & Continue
	for i := 1; i <= 10; i++ {
		if i % 2 == 1 {
			continue
		}
		if i > 8 {
			break
		}
		fmt.Println("Angka", i)
	}
	fmt.Println()

	// Perulangan bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
	fmt.Println()

	// Label dalam perulangan
	outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}

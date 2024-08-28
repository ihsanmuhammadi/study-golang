package main

import "fmt"

func main()  {
	// Non-desimal
	var positiveNumber uint8 = 89
	var negativeNumber = -12345678

	fmt.Println(positiveNumber)
	fmt.Printf("%d", negativeNumber)

	// Desimal
	var decimalNumber1 = 2.62
	decimalNumber2 := 2.676

	fmt.Printf("%f %.2f", decimalNumber1, decimalNumber2)

	// Boolean
	var exist bool = true
	nope := false

	fmt.Printf("%t\n%t\n", exist, nope)

	// String
	var message string = "Halo"
	message1 := `Nama saya adalah
	Ihsan Muhammad Iqbal`

	fmt.Printf("%s %s", message, message1)

	//
}

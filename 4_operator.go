package main

import "fmt"

func main()  {
	// Operator aritmatika
	value := (((2 + 6) % 3) * 4 - 2) / 3
	fmt.Println(value)

	// Operator perbandingan
	var isEqual = (value <= 1)
	fmt.Printf("nilai %t\n", isEqual)

	// Operator logika
	var left bool = true
	right := false

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t)\n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t)\n", leftOrRight)

	var leftReverse = !right
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
}

package main

import (
	"fmt"
)

func main()  {
	// if, else if, else
	var point uint8 = 8

	if point == 10 {
		fmt.Println("Lulus sempurna")
	} else if point > 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("Hampir lulus")
	} else {
		fmt.Printf("Tidak lulus, nilai anda %d", point)
	}

	// Variabel temporary pada if-else
	var point1 = 100

	if percent := point1 / 10; percent >= 10{
		fmt.Println("perfect")
	} else if percent >= 7{
		fmt.Println("good")
	} else {
		fmt.Println("not bad")
	}

	// Switch case
	var point2 = 6

	switch point2 {
	case 8:
		fmt.Println("perfect")
	case 7, 1, 2, 3:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("It's okay")
		}
	}

	// Switch dengan if-else
	var point3 = 6
	switch {
	case point3 == 8:
 		fmt.Println("perfect")
	case (point3 < 8) && (point3 > 3):
 		fmt.Println("awesome")
	default:
 		{
 			fmt.Println("not bad")
 			fmt.Println("you need to learn more")
 		}
	}

	// Switch : penggunaan fallthrough
	var point4 = 7
	switch point4 {
	case 3:
		fmt.Println("Kurang dari 7")
	case 7:
		fmt.Println("Ini tujuh")
		fallthrough
	default :
		fmt.Println("Ini bukan tujuh")
	}

	// If-else bersarang
	var point5 = 10

	if point5 > 7 {
		switch point5 {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point5 == 5 {7
			fmt.Println("not bad")
		} else if point5 == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point5 == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}

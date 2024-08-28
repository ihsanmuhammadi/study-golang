package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

func main()  {
	var names = []string{"John", "Wick"}
	printMessage("Halo", names)
	fmt.Println()

	var randomValue = randomWithRange(2,10)
	fmt.Println(randomValue)
	fmt.Println()

	divideNumber(10, 2)
	divideNumber(4, 0)
	fmt.Println()
}

// Penerapan fungsi (void function)
func printMessage(message string, arr []string)  {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

// Fungsi dengan Return value/nilai balik
func randomWithRange(min, max int) int {
	var value = randomizer.Int()%(max-min+1) + min
	return value
}

// Keyword return untuk menghentikan fungsi
func divideNumber(m, n int)  {
	if n == 0 {
		fmt.Printf("Invalid divider. %d cannot divided by %d\n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}

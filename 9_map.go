package main

import (
	"fmt"
)

func main()  {
	// map
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["Januari"] = 50
	chicken["Februari"] = 40

	fmt.Println("Januari", chicken["Januari"])
	fmt.Println("Mei", chicken["Mei"])
	fmt.Println()

	// Iterasi Map menggunakan for-range
	var bulan = map[string]int{
		"Januari": 50,
		"Februari": 40,
		"Maret": 34,
		"April": 67,
	}
	for key, val := range bulan {
		fmt.Println(key, " \t:", val)
	}
	fmt.Println()

	// Menghapus item Map
	delete(chicken, "Januari")
	fmt.Println(chicken)
	fmt.Println()

	// Deteksi keberadaan Item dengan key
	var value, isExist = chicken["Mei"]
	if isExist {
		fmt.Print(value)
	} else {
		fmt.Println("Item is not exists")
	}
	fmt.Println()

	// Kombinasi slice & map
	var foods = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	for _, food := range foods {
		fmt.Println(food["gender"], food["name"])
	}
	fmt.Println()
}

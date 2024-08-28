package main

import "fmt"

func main()  {
	// Array
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "law"
	names[3] = "water"

	fmt.Println(names[0], names[1], names[2], names[3])
	fmt.Println()

	// Inisialisasi nilai awal Array
	var fruits = [4]string{"Apple", "Grape", "Banana", "Melon"}
	fruits = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}
	fmt.Println("Jumlah elemen \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)
	fmt.Println()

	// Inisialisasi nilai awal array tanpa jumlah elemen
	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))
	fmt.Println()

	// Array multidimensi
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)
	fmt.Println()

	//  Perulangan elemen array menggunakan for
	var students = [4]string{"Hawa", "Ihsan", "Lestari", "Adhan"}

	for i := 0; i < len(students); i++ {
		fmt.Printf("Elemen %d : %s\n", i, students[i])
	}
	fmt.Println()

	// Perulangan array mengguanakan for-range
	for i, student := range students {
		fmt.Printf("For range %d : %s\n", i, student)
	}
	fmt.Println()

	// Penggunaan variabel underscore dalam for-range (jika tidak memerlukan indeks)
	for _, student := range students {
		fmt.Printf("No index : %s\n", student)
	}
	fmt.Println()

		/* Jika nilai tidak digunakan */
	for i, _ := range students{
		fmt.Printf("No value : %d\n", i)
	}
	fmt.Println()

	// Alokasi elemen array dengan keyword "make"
	var company = make([]string, 2)
	company[0] = "Unilever"
	company[1] = "TransTrack"

	fmt.Println(company)
}

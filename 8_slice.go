package main

import "fmt"

func main()  {
	// Inisialisasi slice (jumlah elemen tidak dituliskan)
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0])
	fmt.Println()

	// Hubungan slice dengan array & operasi slice
	var students = [4]string{"Hawa", "Adhan", "Lestari", "Ihsan"}
	var newStudents = students[0:2]
	var allStudents = students[:]
	fmt.Println(newStudents)
	fmt.Println()
	fmt.Println(allStudents)
	fmt.Println()

	// Slice merupakan tipe data reference
	var aFruits = fruits[0:3]
	var aaFruits = fruits[1:2]
	fmt.Println(aFruits)
	fmt.Println(aaFruits)

		/*"grape" menjadi "pineapple"*/
	aaFruits[0] = "Pineapple"
	fmt.Println(aFruits)
	fmt.Println()

	// Fungsi len
	fmt.Println(len(fruits))
	fmt.Println()

	// Fungsi cap
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	var bFruits = fruits[0:3]
	fmt.Println(len(bFruits))
	fmt.Println(cap(bFruits))
	fmt.Println()

	// Fungsi append
	var appendFruits = append(fruits, "papaya")
	fmt.Println(appendFruits)
	fmt.Println()

	// Fungsi copy
	dst := make([]string, 2)
	src := []string{"watermelon", "pineapple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)
	fmt.Println()

	// Copy dengan src <= dst
	dest := []string{"potato", "potato", "potato"}
	source := []string{"watermelon", "pineapple"}
	x := copy(dest, source)

	fmt.Println(dest)
	fmt.Println(source)
	fmt.Println(x)
	fmt.Println()

	// Akses elemen slice dengan 3 index (index ke 3 -> cap)
	var cFruits = fruits[0:2:2]
	fmt.Println(len(cFruits))
	fmt.Println(cap(cFruits))
}

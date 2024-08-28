package main

import "fmt"

func main()  {
	var firstName string = "Ihsan"
	var lastName  string
	lastName = "Iqbal"
	middleName := "Muhammad"
	middleName = "Muhammad"
	var animal = "Sapi"
	_ = "Ini"
	_ = "Tidak akan dipakai"
	keywordNew := new(string)
	fmt.Println(keywordNew)
	fmt.Println(*keywordNew)
	

	multi, variabel := 1, "2"

	fmt.Println("Hello World" + animal)
	fmt.Printf("Halo %s %s!\n", firstName, lastName)
	fmt.Println("Halo", firstName, middleName, lastName + "!")
	fmt.Println(multi, variabel + "Oke")
}

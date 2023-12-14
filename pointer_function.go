package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{}
	ChangeCountryToIndonesia(&address)

	fmt.Println(address) // tidak berubah
}
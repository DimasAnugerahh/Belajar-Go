package main

import "fmt"

type mobil struct {
	merk, Type string
	cc         int
}

func (list mobil) dataMobil(){
	fmt.Println(list)
}

func main() {
	var honda mobil
	var toyota mobil
	var nissan mobil

	honda.Type = "jazz"
	honda.merk = "honda"
	honda.cc = 1500

	// fmt.Println(data)

	toyota = mobil{
		merk: "Toyota",
		Type: "Fortuner",
		cc: 2400,
	}

	// fmt.Println(toyota)

	nissan =mobil{"Nissan","Silvia",3000}
	

	// fmt.Println(nissan)

	//panggil struct method
	honda.dataMobil()
	toyota.dataMobil()
	nissan.dataMobil()

	
}
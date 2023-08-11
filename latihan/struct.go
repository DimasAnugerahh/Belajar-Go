package main

import "fmt"

type mobil struct {
	merk, Type string
	cc         int
}

func main() {
	var data mobil

	data.Type = "jazz"
	data.merk = "honda"
	data.cc = 1500

	fmt.Println(data)


	toyota:=mobil{
		merk: "Toyota",
		Type: "Fortuner",
		cc: 2400,
	}

	fmt.Println(toyota)

	nissan:=mobil{"Nissan","Silvia",3000}
	

	fmt.Println(nissan)

}
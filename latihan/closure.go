package main

import "fmt"

func main() {

	angka:=1 //angka ini tidak dipengaruhi oleh increment pada function

	increment := func() {
		angka:=5 //variable ini dibuat untuk fungsi
		angka++ //increment ini berlaku untuk variable angka pada fungsi
		fmt.Println(angka)
	}

	increment()

	fmt.Println(angka)

}
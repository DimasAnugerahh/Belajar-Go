package main

import "fmt"

func multipleReturn() (string, string) {
	return "Dimas", "Anugerah"
}

func kali() (int,int){
	return 5,4
} 

func main() {

	firstName, lastName := multipleReturn() //assign function value to variable
	namaPertama,_:=multipleReturn() //abaikan salah satu return value menggunakan underscore
	angka1,angka2:=kali()
	fmt.Println(firstName,lastName)
	fmt.Println(angka1*angka2)


	//mengabaikan salah satu return value
fmt.Println(namaPertama)
}
package main

import "fmt"

func nama() (firstName, middleName, lastName string) {
	firstName = "Dimas"
	middleName = " anugerah"
	lastName = " pratama"

	return
}

func main() {
	first, _, _ := nama()

	fmt.Println(first)
	fmt.Println(nama())
}
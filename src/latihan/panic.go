package main

import (
	"fmt"
)

//panic and recover are such a error handling in golang

func logging() {
	fmt.Println("APLIKASI SELESAI")
	message:=recover() //message are from panic message
	//recover same like continue in other programming language
	if message!=nil {
		fmt.Println("Error: ",message) //execute when error
	}
}

func runApp(error bool) {
	defer logging() //async statement that will execute when entire statement are done
	if error {
		panic("There something wrong ") //exception
	}

	fmt.Println("aplikasi berhasil")
}

func main() {
	runApp(false) //kondisi tidak error
	fmt.Println("\n\n")
	runApp(true) //kondisi error 

}
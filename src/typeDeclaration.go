package main

import "fmt"

func main() {

	type noKtp string //noktp digunakan sebagai alias tipe data string

	var kk string ="123" //umum
	var ktp noKtp = "7309132603030001" //implementasi noktp sebagai pengganti string

	fmt.Println(ktp)
	fmt.Print(kk)


}
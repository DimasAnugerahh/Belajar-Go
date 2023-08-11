package main

import "fmt"

func faktorial(angka int) int {
	if angka < 1 {
		return 1
	} else {
		return angka * (faktorial(angka - 1))
	}
}

func main() {
	var angka int;

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)
	fmt.Print("hasil dari ",angka, " faktorial adalah ",faktorial(angka))
}
package main

import "fmt"

func main() {
	var b int
	fmt.Scan(&b)
	kg := b / 1000
	gr := b % 1000
	hargakg := kg * 10000
	hargagr := gr * 15
	if kg > 10 {
		hargagr = 0
	} else if gr >= 500 {
		hargagr = gr * 5
	}
	total := hargakg + hargagr
	fmt.Println("Berat parsel (gram) :", b)
	fmt.Println("Detail berat :", kg, "kg +", gr, "gr")
	fmt.Println("Detail biaya : Rp.", hargakg, "+ Rp.", hargagr)
	fmt.Println("Total biaya : Rp.", total)
}

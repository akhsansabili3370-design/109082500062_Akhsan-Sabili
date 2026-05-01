package main

import "fmt"

func main() {
	var klub1, klub2 string
	var skor1, skor2 int

	fmt.Print("Klub A : ")
	fmt.Scan(&klub1)
	fmt.Print("Klub B : ")
	fmt.Scan(&klub2)

	hasil := []string{}
	for {
		fmt.Scan(&skor1, &skor2)

		if skor1 < 0 || skor2 < 0 {
			break
		}
		if skor1 > skor2 {
			hasil = append(hasil, klub1)
		} else if skor1 < skor2 {
			hasil = append(hasil, klub2)
		} else {
			hasil = append(hasil, "DRAW")
		}
	}
	for i := 0; i < len(hasil); i++ {
		fmt.Println("Hasil ", i+1 ,":", hasil[i])
	}
}

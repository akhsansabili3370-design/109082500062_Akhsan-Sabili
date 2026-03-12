package main

import "fmt"

func main() {
	var a, b, c, d string

	p := [4]string{"merah", "kuning", "hijau", "ungu"}

	hasil := true
	for i := 1; i <= 5; i++ {
		fmt.Scanln(&a, &b, &c, &d)

		if a != p[0] || b != p[1] || c != p[2] || d != p[3] {
			hasil = false
		}
	}
	if hasil == false {
		fmt.Println("False")
	} else {
		fmt.Println("True")
	}
}

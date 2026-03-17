package main

import "fmt"

func faktor(n int) int {
	var hasil = 1
	for i := 2; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutasi(n, r int) int {
	return faktor(n) / faktor(n-r)
}

func kombinasi(n, r int) int {
	return faktor(n) / (faktor(r) * faktor(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(permutasi(a, c), kombinasi(a, c))
	fmt.Println(permutasi(b, d), kombinasi(b, d))
}

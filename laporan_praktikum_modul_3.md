# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">[AKHSAN SABILI] - [109082500062]</p>

## Unguided 

### 1. [Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p)]

#### soal1.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[penjelasan]

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
![Screenshot Output Unguided 1_1](https://github.com/akhsansabili3370-design/109082500062_Akhsan-Sabili/blob/main/P%20modul%203/screenshotm3/sstugas1.png)

[Program diatas digunakan untuk menghitung permutasi dan kombinasi. Pertama terdapat fungsi faktor yang digunakan untuk menghitung faktorial yang juga berfungsi sebagai dasar dari permutasi dan kombinasi. Selanjutnya fungsi permutasi yang berfungsi untuk menghitung banyaknya cara mendapat r dari n. Lalu kombinasi yang berfungsi untuk menghitung banyaknya cara untuk memilih sesuatu berdasarkan faktorialnya. Lalu program akan meminta empat inputan dengan tipe interger, kemudian akan ditampilkan hasil perhitungan permutasi dan kombinasinya]

### 2. [Diberikan tiga buah fungsi matematika yaitu f (x) = x^2, g (x) = x − 2 dan h (x) = x + 1. Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x) dalam bentuk function.]

#### soal2.go

```go
package main

import "fmt"

func f(x int) int {
	return x * x
}
func h(x int) int {
	return x + 1
}
func g(x int) int {
	return x - 2
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(f(g(h(a))))
	fmt.Println(g(h(f(b))))
	fmt.Println(h(f(g(c))))
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/akhsansabili3370-design/109082500062_Akhsan-Sabili/blob/main/P%20modul%203/screenshotm3/sstugas2.png)

[Program diatas berfungsi untuk mendefinisikan fungsi. Pertama f(x) yang digunakan untuk menghitung kuadrat, h(x) digunakan untuk menambah x dengan 1, dan g(x) digunakan untuk mengurangi x dengan 2. Pada func main program akan membaca tiga inputan pengguna dengan tipe interger. Kemudian program akan menampilkan output berupa (fogoh)(x), (gohof)(x) dan (hofog)(x)]

### 3. [ Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut.]

#### soal3.go

```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}
func dalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}
func main() {
	var a, b, c, d, e, f float64
	var x, y float64

	fmt.Scan(&a, &b, &c)
	fmt.Scan(&d, &e, &f)
	fmt.Scan(&x, &y)
	bulat := dalam(a, b, c, x, y)
	bulat2 := dalam(d, e, f, x, y)

	if bulat && bulat2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if bulat {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if bulat2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/akhsansabili3370-design/109082500062_Akhsan-Sabili/blob/main/P%20modul%203/screenshotm3/sstugas3.png)

[Program diatas berfungsi untuk menentukan apakah sebuah titik berada di dalam satu atau dua lingkaran. Pertama terdapat fungsi jarak untuk menghitung jarak antara dua titik, lalu fungsi dalam yang berfungsi untuk memeriksa apakah titik uji berada di dalam lingkaran dengan membandingkan jarak ke pusat dengan r atau jari jari. Berikutnya func main yang digunakan untuk membaca inputan berupa pusat dan jari-jari dua lingkaran serta koordinat titik. Lalu mengecek posisi titik terhadap lingkaran. Berikutnya terdapat empat if else untuk menentukan apakah titik berada di dalam lingkaran 1, lingkaran 2, keduanya, atau di luar lingkaran]

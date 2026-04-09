# <h1 align="center">Laporan Praktikum Modul 5 </h1>
<p align="center">[AKHSAN SABILI] - [109082500062]</p>

## Unguided 

### 1. [Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.]

#### soal1.go

```go
package main
import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	for i := 0; i <= n; i++{
		fmt.Print(fibo(i), " ")
	}
}
func fibo(n int)int{
	if n <= 1{
		return n
	}
	return fibo(n - 1) + fibo(n - 2)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/akhsansabili3370-design/109082500062_Akhsan-Sabili/blob/main/P%20modul%205/screenshot/ss_tugas1.png)

[Program diatas berfungsi untuk menampilkan deret Fibonacci mulai dari 0 hingga angka yang dimasukkan pengguna.pertama program akan membaca input n lalu melakukan perulangan dari 0 sampai n. Pada setiap i akan memanggil fungsi fibo(i) untuk menghitung fibonacci ke i. Fungsi fibo menggunakan rekursi jika n bernilai 0 atau 1 akan kembali ke n dasn jika lebih besar akan dihitung menjadi fibo(n-1) + fibo(n-2).]


### 2. [Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.]

#### soal2.go

```go
package main
import "fmt"

func bintang(n int) {
	if n != 0{
		fmt.Print("*")
		bintang(n - 1)
	}
}
func main(){
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++{
		bintang(i)
		fmt.Println()
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/akhsansabili3370-design/109082500062_Akhsan-Sabili/blob/main/P%20modul%205/screenshot/ss_tugas2.png)

[Program diatas berfungsi untuk mencetak segitiga dengan * sesuai input pengguna. Func bintang berfungsi untuk mencetak bintang sebanyak n menggunakan rekursi yang jika n tidak sama dengan 0 maka akan tercetak satu bintang lalu memanggil dirinya sendiri dengan n - 1. Pada func main. program membaca input angka n lalu melakukan perulangan dari 1 hingga n. Setiap i memanggil bintang(i) untuk mencetak * sebanyak i, kemudian pindah ke baris baru]


### 3. [Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N.]

#### soal3.go

```go
package main
import "fmt"
func faktor( n, m int){
	if n % m == 0{
		fmt.Print(m, " ")
	}
	if n > m{
		faktor(n, m + 1)
	}
}
func main(){
	var n int
	fmt.Scan(&n)
	faktor(n, 1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/akhsansabili3370-design/109082500062_Akhsan-Sabili/blob/main/P%20modul%205/screenshot/ss_tugas3.png)

[Program diatas berfungsi untuk menampilkan semua faktor sesuai inputan pengguna. Program membaca input n lalu memanggil func faktor. Fungsi ini bekerja secara rekursif dengan memeriksa dari 1 hingga n. Jika n habis dibagi oleh m, maka m dicetak sebagai faktor. jika m lebih kecil dari n, fungsi akan memanggil dirinya dengan nilai m+1. Lalu program akan menampilkan daftar faktor dari bilangan n yang ditampilkan berurutan]


### 4. [Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan tertentu.]

#### soal4.go

```go
package main
import "fmt"
func baris(n int){
	for i := n; i >= 2; i--{
		fmt.Print(i," ")
	}
	for j := 1; j <= n; j++{
		fmt.Print( j, " ")
	}
}
func main(){
	var n int
	fmt.Scan(&n)
	baris(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/akhsansabili3370-design/109082500062_Akhsan-Sabili/blob/main/P%20modul%205/screenshot/ss_tugas4.png)

[Program diatas berfungsi untuk mencetak deretan angka dengan pola dari n ke 1 dan kembali lagi ke n. Program akan membaca sebuah bilangan n lalu memanggil func baris. Di dalam fungsi tersebut terdapat dua perulangan yaitu perulangan pertama berjalan dari n turun hingga angka 2, sehingga mencetak angka menurun mulai dari n sampai 2 dan perulangan kedua berjalan dari 1 naik hingga n, sehingga mencetak angka naik dari 1 sampai n.]

### 5. [Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil.]

#### soal5.go

```go
package main
import "fmt"
func ganjil( n int){
	for i := 1; i <= n; i++{
		if i % 2 != 0{
			fmt.Print(i," ")
		}
	}
}
func main(){
	var n int
	fmt.Scan(&n)
	ganjil(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/akhsansabili3370-design/109082500062_Akhsan-Sabili/blob/main/P%20modul%205/screenshot/ss_tugas5.png)

[Program diatas berfungsi untuk menampilkan bilangan ganjil dari 1 hingga n. Program akan membaca input n lalu memanggil func ganjil. Di dalam func ganjil terdapat perulangan dari 1 sampai n. Setiap angka diperiksa dengan kondisi if i % 2 != 0 yang berarti hanya angka yang tidak habis dibagi 2 yang akan dicetak. Hasil akhirnya adalah deret bilangan ganjil dari 1 sampai n]

### 6. [Buatlah program yang mengimplementasikan rekursif untuk mencari hasil pangkat dari dua buah bilangan.]

#### soal6.go

```go
package main
import ("fmt"
		"math")
func pangkat(n, m int){
	fn := float64(n)
	fm := float64(m)
	hasil := math.Pow(fn, fm)
	fmt.Println(hasil)
}
func main(){
	var x, y int
	fmt.Scan(&x, &y)
	pangkat(x, y)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/akhsansabili3370-design/109082500062_Akhsan-Sabili/blob/main/P%20modul%205/screenshot/ss_tugas6.png)

[Program diatas berfungsi untuk menghitung hasil perpangkatan dari sebuah bilangan. Program membaca dua input integer dari pengguna, yaitu x sebagai bilangan dan y sebagai pangkat yang kemudian kirim ke func pangkat. Di dalam fungsi, kedua nilai diubah menjadi float64 agar bisa diproses oleh math.Pow dari paket math. Lalu output akan menampilkan hasil dari perpangkatan tersebut]
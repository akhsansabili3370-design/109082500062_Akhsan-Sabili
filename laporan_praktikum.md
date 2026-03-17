# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">[Akhsan Sabili] - [109082500062]</p>

## Unguided 

### 1. [Soal]
#### tugas1.go

```go
package main
import "fmt"
func main(){
	var (
		satu, dua, tiga string
		temp string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/akhsansabili3370-design/109082500062_Akhsan-Sabili/blob/main/screenshot/ss%20tugas%201.png)

[Program tersebut meminta tiga inputan string dari pengguna, lalu program akan menampilkan hasil awal berupa urutan sesuai yang di input pengguna. Setelah menampilkan output awal, program akan melakukan pergeseran ke kiri dari urutan output awal. Berikutnya program akan menampilkan output kedua berupa hasil dari pergeseran ke kiri. Contohnya : jika pengguna menginput A B C maka hasil akhirnya menjadi B C A.]

### 2. [Soal]
#### tugas2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/akhsansabili3370-design/109082500062_Akhsan-Sabili/blob/main/screenshot/ss%20tugas%202.png)

[Program diatas memiliki empat variable utama dengan tiper string. Lalu terdapat variable p untuk menyimpan empat string secara berurutan dan variable hasil dengan nilai inisiasi awal true. Berikutnya terdapat perulangan yang akan menjalankan perulangan sebanyak lima kali. Dalam looping terdapat Scan yang berfungsi untuk membuat inputan berulang sebanyak lima kali. Lalu terdapat if yang bermaksud menyamakan apakah a = merah, b = kuning, c = hijau, d = ungu, jika diantara warna tersebut ada yang tidak sesuai maka variable hasil akan berubah menjadi false. Lalu terdapat dua output yang apabila hasil sama dengan false maka program akan menampilkan output "False" dan jika hasil sama dengan true maka program akan menampilkan  output "True"]


### 3. [Soal]
#### tugas3.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/akhsansabili3370-design/109082500062_Akhsan-Sabili/blob/main/screenshot/ss%20tugas%203.png)

[Program diatas memiliki satu variable utama dengan tipe interger. Lalu terdapat beberapa variable tambahan seperti kg untuk mengonversi dari gram ke kg, lalu gram digunakan untuk menghitung sisa dari kg, lalu hargakg dan hargagr digunakan untuk menentukan harga perkg dan per gram. Berikutnya terdapat if yang bermaksud jika berat lebih dari 10 kg maka harga pergram menjadi 0, lalu else if yang bermaksud jika sisa berat barang lebih dari atau sama dengan 500 gram maka harga pergram menjadi 5 rupiah. Lalu variable total berfungsi untuk menjumlahkan harga akhir dari kg dan gram. Setelah semua selesai maka output akan menampilkan mulai dari berat parsel detail berat, detail harga dan total biaya]

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

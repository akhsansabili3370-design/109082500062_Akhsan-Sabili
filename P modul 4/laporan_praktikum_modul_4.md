# <h1 align="center">Laporan Praktikum Modul 4 </h1>
<p align="center">[AKHSAN SABILI] - [109082500062]</p>

## Unguided 

### 1. [Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas?]

#### soal1.go

```go
package main
import "fmt"
func factorial(n int, hasil *int64) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= int64(i)
	}
}
func permutation(n, r int, hasil *int64) {
	var fn, fnr int64
	factorial(n, &fn)
	factorial(n-r, &fnr)
	*hasil = fn / fnr
}
func combination(n, r int, hasil *int64) {
	var fn, fr, fnr int64
	factorial(n, &fn)
	factorial(r, &fr)
	factorial(n-r, &fnr)
	*hasil = fn / (fr * fnr)
}
func main() {
	var a, b, c, d int
	var x, y int64
	fmt.Scan(&a, &b, &c, &d)
	permutation(a, c, &x)
	combination(a, c, &y)
	fmt.Println(x, y)
	permutation(b, d, &x)
	combination(b, d, &y)
	fmt.Println(x, y)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/akhsansabili3370-design/109082500062_Akhsan-Sabili/blob/main/P%20modul%204/screenshot_modul_4/Screenshot%20tugas%201.png)

[Kode diatas digunakan untuk menghitung nilai permutasi dan kombinasi dari dua pasang bilangan sesuai dengan inputan. Cara kerjanya ialah ada fungsi factorial yang menghitung faktorial suatu bilangan dengan mengalikan angka dari 1 sampai n. Fungsi ini kemudian dipakai oleh permutasi untuk menghitung banyaknya susunan berbeda dari n objek yang diambil r objek, dan oleh kombinasi untuk menghitung banyaknya cara memilih r objek dari n tanpa memperhatikan urutan. Selanjutnya pada func main program membaca empat angka dari input, lalu menampilkan hasil permutasi dan kombinasi untuk pasangan pertama (a, c) dan pasangan kedua (b, d)]

### 2. [Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur.]

#### soal2.go

```go
package main
import "fmt"

func hitungSkor(waktu1, waktu2, waktu3, waktu4, waktu5, waktu6, waktu7, waktu8 int, soal *int, skor *int) {
    *soal = 0
    *skor = 0
    if waktu1 <= 300 {
        *soal++
        *skor += waktu1
    }
    if waktu2 <= 300 {
        *soal++
        *skor += waktu2
    }
    if waktu3 <= 300 {
        *soal++
        *skor += waktu3
    }
    if waktu4 <= 300 {
        *soal++
        *skor += waktu4
    }
    if waktu5 <= 300 {
        *soal++
        *skor += waktu5
    }
    if waktu6 <= 300 {
        *soal++
        *skor += waktu6
    }
    if waktu7 <= 300 {
        *soal++
        *skor += waktu7
    }
    if waktu8 <= 300 {
        *soal++
        *skor += waktu8
    }
}

func main() {
    var nama, pemenang string
    var soal, skor int
    var maxSoal, minSkor int
    var first = true

    for {
        fmt.Scan(&nama)
        if nama == "Selesai" {
            break
        }

        var w1, w2, w3, w4, w5, w6, w7, w8 int
        fmt.Scan(&w1, &w2, &w3, &w4, &w5, &w6, &w7, &w8)

        hitungSkor(w1, w2, w3, w4, w5, w6, w7, w8, &soal, &skor)

        if first || soal > maxSoal || (soal == maxSoal && skor < minSkor) {
            first = false
            maxSoal = soal
            minSkor = skor
            pemenang = nama
        }
    }

    fmt.Println(pemenang, maxSoal, minSkor)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/akhsansabili3370-design/109082500062_Akhsan-Sabili/blob/main/P%20modul%204/screenshot_modul_4/Screenshot%20tugas%202.png)

[Program ini intinya mencari siapa peserta yang paling unggul dalam lomba. Setiap orang memasukkan nama dan waktu pengerjaan delapan soal. Kalau waktu sebuah soal tidak lebih dari 300 detik, soal itu dihitung berhasil dan waktunya ditambahkan ke skor. Setelah semua peserta dimasukkan, program memilih pemenang berdasarkan jumlah soal terbanyak, dan jika sama, dilihat dari total waktu paling sedikit. Hasil akhirnya menampilkan nama pemenang, jumlah soal yang berhasil, serta total waktunya.]

### 3. [Skiena dan Revilla dalam Programming Challenges mendefinisikan sebuah deret bilangan. Deret dimulai dengan sebuah bilangan bulat n. Jika bilangan n saat itu genap, maka suku berikutnya adalah 1⁄2n, tetapi jika ganjil maka suku berikutnya bernilai 3n+1. Rumus yang sama digunakan terus menerus untuk mencari suku berikutnya. Deret berakhir ketika suku terakhir. Halaman 48 | M o d u l P r a k t i k u m A l g o r i t m a d a n P e m r o g r a m a n 2 bernilai 1. Sebagai contoh jika dimulai dengan n=22, maka deret bilangan yang diperoleh adalah: 22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1 Untuk suku awal sampai dengan 1000000, diketahui deret selalu mencapai suku dengan nilai 1.Buat program skiena yang akan mencetak setiap suku dari deret yang dijelaskan di atas untuk nilai suku awal yang diberikan. Pencetakan deret harus dibuat dalam prosedur cetakDeret yang mempunyai 1 parameter formal, yaitu nilai dari suku awal]
#### soal3.go

```go
package main
import "fmt"
func cetakDeret(p int) {
	for {
		fmt.Print(p," ")

		if p == 1 {
			break
		}
		if p % 2 == 0 {
			p = p / 2
		} else {
			p = 3 * p + 1
		}
	}
}

func main() {
	var x int
	fmt.Scan(&x)
	cetakDeret(x)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/akhsansabili3370-design/109082500062_Akhsan-Sabili/blob/main/P%20modul%204/screenshot_modul_4/Screenshot%20tugas%203.png)

[Program ini berfungsi untuk mencetak deret angka berdasarkan aturan yang dikenal sebagai Collatz sequence. Cara kerjanya adalah: pengguna memasukkan sebuah bilangan, kemudian program akan menampilkan angka tersebut. Jika angka yang dimasukkan bernilai 1, proses berhenti. Namun, jika angka tersebut genap, maka akan dibagi dua, sedangkan jika angka ganjil, angka tersebut dikalikan tiga lalu ditambah satu. Hasil perhitungan ini kemudian menjadi angka baru yang kembali diproses dengan aturan yang sama, hingga akhirnya mencapai angka 1.]


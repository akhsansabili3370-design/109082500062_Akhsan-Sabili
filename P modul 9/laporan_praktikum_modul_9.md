# <h1 align="center">Laporan Praktikum Modul 7 & 9 </h1>
<p align="center">[AKHSAN SABILI] - [109082500062]</p>

## Unguided 

### 1. [Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya.]

#### soal1.go

```go
package main
import "fmt"

func dalamLingkaran(cx, cy, r, x, y int) bool {
	dx := x - cx
	dy := y - cy
	return dx * dx + dy * dy <= r * r
}

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	dalam1 := dalamLingkaran(cx1, cy1, r1, x, y)
	dalam2 := dalamLingkaran(cx2, cy2, r2, x, y)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1]()
[penjelasan]

# <h1 align="center">Laporan Praktikum Modul 5 </h1>
<p align="center">[AKHSAN SABILI] - [109082500062]</p>

## Unguided 

### 1. [Soal]
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
![Screenshot Output Unguided 1_1]()
[penjelasan]

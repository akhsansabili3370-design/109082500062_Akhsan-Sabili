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
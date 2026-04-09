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
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
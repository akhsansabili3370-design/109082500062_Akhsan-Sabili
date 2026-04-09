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
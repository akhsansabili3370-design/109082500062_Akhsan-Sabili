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
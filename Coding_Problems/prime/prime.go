package main
import (
	"fmt"
	"math"
)

func is_prime(num int) bool{
	if num < 2{
		return false
	}
	for i:=2;i<=int(math.Sqrt(float64(num)));i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(is_prime(67))
	fmt.Println(is_prime(45))
	fmt.Println(is_prime(81))
}
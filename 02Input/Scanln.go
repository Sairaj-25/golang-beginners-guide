package main
import "fmt"

func main() {
	var fullname string
	fmt.Print("Enter your fullname:")
	fmt.Scanln(&fullname)
	fmt.Println("Hello,",fullname)
}

// Issue: fmt.Scanln only captures the first word if multiple words are entered.
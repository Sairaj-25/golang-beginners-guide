package main
import "fmt"

func main() {
	var marks int
	fmt.Print("Enter Your marks:")
	_, err := fmt.Scanf("%d", &marks) // Ensures only integers are accepted

	if err != nil {
		fmt.Println("Please enter a valid input !")
		return
	}

	if marks >= 90 {
		fmt.Println("Grade A")
	} else if marks>=75 {
		fmt.Println("Grade B")
	} else if marks >= 50 {
		fmt.Println("Grade c")
	} else if marks>=40 {
		fmt.Println("Grade d")
	} else {
		fmt.Println("Grade F")
	}
}
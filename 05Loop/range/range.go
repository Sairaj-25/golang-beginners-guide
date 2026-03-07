package main
import "fmt"
func main() {
	/*
	numbers := []int{10, 20, 30, 40, 50}

	for index, value := range numbers {
		fmt.Printf("Index: %d, value: %d\n",index,value)
	}
	*/

	str := "Golang"
	for index, char := range str {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

}
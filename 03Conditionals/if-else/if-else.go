package main

import ("fmt")

func main(){
	var num int
	fmt.Print("Enter a number:")
	_, err := fmt.Scan(&num)

	if err != nil {
		fmt.Print("Invalid input, please enter a valid number!")
		return
	}

	if num%2 == 0 {
		fmt.Println("The number is Even")
	} else {
		fmt.Println("The number is Odd")
	}
}
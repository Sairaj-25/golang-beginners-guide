package main

import "fmt"


// In one package or one folder, Go allows only ONE main() function.

// You currently have two main() functions in the same package (usually package main).

func main(){
	for i:= 1; i<=5; i++ {
		fmt.Println("Iteration",i)
	}
}
package main

// import "fmt"


func add(a int, b int) int {
	return a + b
}

func getLanguages() (string, string, string){
	return "golang", "python", "java"
}


// Variadic Functions

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}
	return total
}
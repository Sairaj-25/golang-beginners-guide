package main

import "fmt"

func main() {
	result := add(3, 5)
	fmt.Println("sum is", result)

	lang1, lang2, lang3 := getLanguages()
	fmt.Println(lang1, lang2, lang3)

	fmt.Println(sum(1, 2, 3, 4))
}


// Use "go run ." to run all files in a folder .
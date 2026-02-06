package main

import "fmt"

func reverseString(s string) string {
	runes := []rune(s) // Convert string to rune slice (to handle Unicode)
	n := len(runes)
	for i:=0; i<n/2;i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func main(){
	str:= "Sairaj"
	fmt.Println(reverseString(str))
}
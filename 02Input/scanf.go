package main

import "fmt"

func main(){
	var age int
	var name string
	fmt.Print("Enter your name and age:")
	fmt.Scanf("%s %d",&name,&age)
	fmt.Printf("Hello %s You are %d years old.\n",name,age)
}
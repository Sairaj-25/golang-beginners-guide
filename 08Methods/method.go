package main

import "fmt"

func main(){
	fmt.Println("Structs in golang")

	sairaj := User{"Sairaj", "sairaj@go.dev", true, 25}
	fmt.Println(sairaj)
	fmt.Printf("sairaj details are: %+v\n", sairaj)
	fmt.Printf("Name is %v and email is %v\n", sairaj.Name, sairaj.Email)

	sairaj.GetStatus()
}

type User struct {
	Name string
	Email string
	status bool
	Age int

}


func (u User) GetStatus(){
	fmt.Println("Is user active: ",u.status)
}
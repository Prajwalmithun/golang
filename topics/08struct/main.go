package main

import (
	"fmt"
	"strings"
)

// create struct
type User struct{
	Name string 
	Email string
	isVerified bool 
}

func main() {
	fmt.Println("Welcome to structs in Golang!!")

	// create a user using struct
	prajwal := User{
		"Prajwal T",
		"prajwal3498@gmail.com",
		false,
	}

	fmt.Printf("Details are %v\n",prajwal)
	fmt.Printf("Details are %+v\n",prajwal)  //more verbose

	fmt.Println(strings.Repeat("#",20))

	// Access each parts of the user
	fmt.Printf("Hi %v!!\n",prajwal.Name) 

	fmt.Println(strings.Repeat("#",20))


	// Create slice of struct 
	usersList := []User{}

	// get details of 5 users
	for i := 0; i < 5; i++ {

		var name, email string 

		fmt.Printf("Enter the details of user %v\n",i+1)
		fmt.Printf("Username: %v",name)
		fmt.Scan(&name)

		fmt.Printf("Email: %v",email)
		fmt.Scan(&email)

		// create a user
		tmpUser := User{}

		tmpUser.Name = name 
		tmpUser.Email = email
		tmpUser.isVerified = false
		
		// append this user to slice
		usersList = append(usersList,tmpUser)
	}

	fmt.Printf("\nHere are the details of the 5 users: %v\n",usersList)

}

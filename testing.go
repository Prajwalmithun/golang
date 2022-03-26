package main

import "fmt"

func main(){

	// // Maps - only stores same datatypes
	// var userData = make(map[string]string) 

	// userData["firstName"] = "Prajwal"
	// userData["lastName"] = "T"
	// userData["email"] = "prajwal@yahoo.com"

	// for key, value := range userData {
	
	// 	fmt.Printf("%v : %v\n", key,value)
		
	// }

	// creating a struct - basically creating custom data type
	type UserData struct {
		 firstName string 
		 lastName string 
		 email string
		 numberOfTickets uint
	}
	
	// creating a variable of struct type
	var userData UserData


	// assigning values to struct items
	userData.firstName = "Prajwal"
	userData.lastName = "t"
	userData.email = "tkushal@gmail.com"
	userData.numberOfTickets = 10

	// printing the values
	var message = fmt.Sprintf("Hi %v!!,\n\nThank you for booking %v tickets.\nWe will also email you the tickets to email address %v \n", userData.firstName, userData.numberOfTickets, userData.email )
	fmt.Printf(message)
}



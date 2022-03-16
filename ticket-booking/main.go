/*
	This code is for "Conference Ticket Booking Application"
*/

package main   //"main" just a standard way, but it can be anything

import "fmt"  //fmt is builtin package for formating, like displaying output to std output 

// entry point to the complier
func main()  {
	var conferenceName = "Golang Conference"    // to store the name of the conference

	// another way to assign a value to variable, just variable NOT contansts
	// conferenceName := "Golang Conference"

	const totalConferenceTickets = 50           // total number of tickets, this is constant overall in conference
	var remainingTickets uint = 50                   // TO store remaining tickets. 

	var firstName string
	var lastName string
	var email string                           // get the user name of the user
	var userTickets int							// get how many tickets user needs

	// we can get the data type of the variable with %T in Printf.
	// fmt.Printf("conferenceName is %T, totalConferenceTickets is %T, username is %T \n", conferenceName,totalConferenceTickets,userName)

	fmt.Print("Welcome to", conferenceName, "booking application \n")
	fmt.Printf("We have total of %v tickets and remaining tickets are %v \n", totalConferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// userName = "John"
	// userTickets = 5

	// fmt.Println(remainingTickets)        // returns the value
	// fmt.Println(&remainingTickets)       // returns the memory address of the value  
	fmt.Println("Enter your first Name")
	fmt.Scan(&firstName)                 // get the input from the users.

	fmt.Println("Enter your last Name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets")
	fmt.Scan(&userTickets)

	// logic of this app
	remainingTickets = remainingTickets - uint(userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation to your email %v\n", firstName,lastName,userTickets,email)
	fmt.Printf("%v remaining tickets.\n",remainingTickets)

}
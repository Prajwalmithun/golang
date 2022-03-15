/*
	This code is for "Conference Ticket Booking Application"
*/

package main   //"main" just a standard way, but it can be anything

import "fmt"  //fmt is builtin package for formating, like displaying output to std output 

// entry point to the complier
func main()  {
	var conferenceName = "Golang Conference"    // to store the name of the conference
	const totalConferenceTickets = 50           // total number of tickets, this is constant overall in conference
	var remainingTickets = 50                   // TO store remaining tickets. 

	var userName string                         // get the user name of the user
	var userTickets int							// get how many tickets user needs

	// we can get the data type of the variable with %T in Printf.
	fmt.Printf("conferenceName is %T, totalConferenceTickets is %T, username is %T \n", conferenceName,totalConferenceTickets,userName)

	fmt.Print("Welcome to", conferenceName, "booking application \n")
	fmt.Printf("We have total of %v tickets and remaining tickets are %v \n", totalConferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend")

	userName = "John"
	userTickets = 5
	fmt.Printf("Assuming the name of the user is %v \n", userName)
	fmt.Printf("Assuming this user need %v tickets \n", userTickets)
	

}
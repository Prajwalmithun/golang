/*
	This code is for "Conference Ticket Booking Application"
*/

package main //"main" just a standard way, but it can be anything

import (
	"fmt"     //fmt is builtin package for formating, like displaying output to std output
	"strings"
	"booking-app/helper" // strings is builting package for handling strings
)

// package level variables
var conferenceName = "Golang Conference" // to store the name of the conference

// another way to assign a value to variable, just variable NOT contansts
// conferenceName := "Golang Conference"

const totalConferenceTickets = 50 // total number of tickets, this is constant overall in conference
var remainingTickets uint = 50    // TO store remaining tickets.

// we can get the data type of the variable with %T in Printf.
// fmt.Printf("conferenceName is %T, totalConferenceTickets is %T, username is %T \n", conferenceName,totalConferenceTickets,userName)

//var bookings = [50]string {"prajwal", "kushal"}     				// arrays in golang

// var bookings [50]string                      					// array to store the names of the users who booked the ticket
var bookings []string // "slices" -> dynamic array
// OR
// var bookings1 = []string{}
// OR
// bookings2 := []string{}

// entry point to the complier
func main() {

	greetUsers()

	/*
		infinte for loop to get the user details.
	*/
	for remainingTickets != 0 {
		firstName, lastName, email, userTickets := getUserDetails()
		isValidName, isValidEmail, isValidUserTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidUserTickets && isValidEmail && isValidName {

			// call bookTicket function
			bookTicket(firstName, lastName, email, userTickets)

			// call printFirstNames function
			printFirstNames()

		} else if userTickets > totalConferenceTickets || userTickets > int(remainingTickets) {
			fmt.Printf("Sorry:( We only have %v tickets remaining. \n",remainingTickets)

		} else {
			//fmt.Printf("We only have %v of total tickets and %v are remaining.\n",totalConferenceTickets,remainingTickets)

			if !isValidName {
				fmt.Printf("First name or last name is too short\n")
			}

			if !isValidEmail {
				fmt.Printf("email address must contain @ and .\n")
			}

			if !isValidUserTickets {
				fmt.Printf("Number of tickets entered are invalid\n")
			}
		}
	}

	if remainingTickets == 0 {
		fmt.Printf("Tickets are sold out, see you next year\n")
	}
}

func greetUsers() {
	fmt.Print("Welcome to ", conferenceName, " booking application \n")
	fmt.Printf("We have a total of %v tickets remaining with us, book faster.\n",remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getUserDetails() (string, string, string, int) {
	var firstName, lastName, email string
	//var lastName string
	//var email string    // get the user name of the user
	var userTickets int // get how many tickets user needs

	fmt.Println("Enter your first Name")
	fmt.Scan(&firstName) // get the input from the users.

	fmt.Println("Enter your last Name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func printFirstNames() {
	// display only the firstName of the users
	var firstNames []string

	//for index, fullName := range bookings{
	for _, fullName := range bookings { // "_" blank identifier, ignore a variable that is unused

		var fullNames = strings.Fields(fullName)
		firstNames = append(firstNames, fullNames[0])
	}
	fmt.Printf("The first names of bookings %v\n", firstNames)
	fmt.Printf("These are all our bookings %v\n", bookings)
}


func bookTicket(firstName string, lastName string, email string, userTickets int) {
	// bookings[0] = firstName + " " + lastName              			// adding elements in "array"
	bookings = append(bookings, firstName+" "+lastName) // adding elements to "slices"

	// logic of this app
	remainingTickets = remainingTickets - uint(userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation to your email %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v remaining tickets.\n", remainingTickets)
}

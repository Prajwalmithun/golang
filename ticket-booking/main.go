/*
	This code is for "Conference Ticket Booking Application"
*/

package main //"main" just a standard way, but it can be anything

import (
	"fmt" //fmt is builtin package for formating, like displaying output to std output
	"time"
	"sync"
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
var bookings = make([]UserData,0) 
// OR
// var bookings1 = make([]map[string]string)
// OR
// bookings2 := map([]map[string]string)


type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint

}

// Creating a wait group for sync main thread with all goroutines.
var wg = sync.WaitGroup{}

// entry point to the complier
func main() {

	greetUsers()

	/*
		infinte for loop to get the user details.
	*/
	for remainingTickets != 0 {
		firstName, lastName, email, userTickets := getUserDetails()
		isValidName, isValidEmail, isValidUserTickets := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidUserTickets && isValidEmail && isValidName {

			// call bookTicket function
			bookTicket(firstName, lastName, email, userTickets)
			

			wg.Add(1)

			// goroutine -> for running the ticket generation and sending tickets in background
			go sendTicket(userTickets,firstName,lastName,email)

			// call printFirstNames function
			printFirstNames()

		} else if userTickets > totalConferenceTickets || userTickets > remainingTickets {
			fmt.Printf("Sorry:( We only have %v tickets remaining. \n", remainingTickets)

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
	wg.Wait()
}

func greetUsers() {
	fmt.Print("Welcome to ", conferenceName, " booking application \n")
	fmt.Printf("We have a total of %v tickets remaining with us, book faster.\n", remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getUserDetails() (string, string, string, uint) {
	var firstName, lastName, email string
	//var lastName string
	//var email string    // get the user name of the user
	var userTickets uint // get how many tickets user needs

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
	for _, booking := range bookings { // "_" blank identifier, ignore a variable that is unused
		
		firstNames = append(firstNames,booking.firstName)
	}
	fmt.Printf("The first names of bookings %v\n", firstNames)
	fmt.Printf("These are all our bookings %v\n", bookings)
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {

	// logic of this app
	remainingTickets = remainingTickets - uint(userTickets)

	// create a empty map to store user's details
	// var userData = make(map[string]string)      // make() : to create empty map
	
	// creating from struct
	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) 

	bookings = append(bookings, userData) // adding elements to "slices"

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation to your email %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v remaining tickets.\n", remainingTickets)
}

// this function is for generating tickets, sending the generated tickets to customers.
// As we see, this takes a while to process the above tasks.
func sendTicket(userTickets uint, firstName string, lastName string, email string){
	var ticket = fmt.Sprintf("%v tickets  for %v %v ", userTickets, firstName, lastName)

	// adding a delay for tickets generation and sending
	fmt.Printf("Generating Ticket...Please wait...\n") 
	time.Sleep(5 * time.Second)        // 5sec delay

	fmt.Printf("Sending the Tickets...\n")
	time.Sleep(5 * time.Second)   // 5 sec delay

	fmt.Println("###########################")
	fmt.Printf("Sent %v to email address: %v\n",ticket,email)
	fmt.Println("###########################")

	wg.Done()
	}	
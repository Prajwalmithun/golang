/*
	This code is for "Conference Ticket Booking Application"
*/

package main 															//"main" just a standard way, but it can be anything

import (
	"fmt" 																//fmt is builtin package for formating, like displaying output to std output
	"strings"															// strings is builting package for handling strings
)

// entry point to the complier
func main()  {
	var conferenceName = "Golang Conference"    						// to store the name of the conference

	// another way to assign a value to variable, just variable NOT contansts
	// conferenceName := "Golang Conference"

	const totalConferenceTickets = 50           						// total number of tickets, this is constant overall in conference
	var remainingTickets uint = 50                   					// TO store remaining tickets. 

	var firstName string
	var lastName string
	var email string                           							// get the user name of the user
	var userTickets int													// get how many tickets user needs

	// we can get the data type of the variable with %T in Printf.
	// fmt.Printf("conferenceName is %T, totalConferenceTickets is %T, username is %T \n", conferenceName,totalConferenceTickets,userName)

	//var bookings = [50]string {"prajwal", "kushal"}     				// arrays in golang

	// var bookings [50]string                      					// array to store the names of the users who booked the ticket
	var bookings []string                            					// "slices" -> dynamic array 
	// OR
	// var bookings1 = []string{}
	// OR
	// bookings2 := []string{} 

	greetUsers(conferenceName, totalConferenceTickets, remainingTickets)

	// userName = "John"
	// userTickets = 5

	// fmt.Println(remainingTickets)        							// returns the value
	// fmt.Println(&remainingTickets)       							// returns the memory address of the value  

		/*
			infinte for loop to get the user details. 
		*/
		for remainingTickets != 0 {
				fmt.Println("Enter your first Name")
				fmt.Scan(&firstName)                 								// get the input from the users.
	
				fmt.Println("Enter your last Name")
				fmt.Scan(&lastName)
	
				fmt.Println("Enter your email address")
				fmt.Scan(&email)
			
				fmt.Println("Enter the number of tickets")
				fmt.Scan(&userTickets)


				var isValidName bool = len(firstName) > 2 && len(lastName) > 2
				var isValidEmail bool = strings.Contains(email,"@")
				var isValidUserTickets bool = userTickets > 0 && userTickets <= int(remainingTickets)

				if isValidUserTickets && isValidEmail && isValidName {

					// bookings[0] = firstName + " " + lastName              			// adding elements in "array"
					bookings = append(bookings, firstName + " " + lastName)  			// adding elements to "slices"
	
					// logic of this app
					remainingTickets = remainingTickets - uint(userTickets)
	
					fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation to your email %v\n", firstName,lastName,userTickets,email)
					fmt.Printf("%v remaining tickets.\n",remainingTickets)
	
					// display only the firstName of the users
					var firstNames []string 
			
					//for index, fullName := range bookings{
					for _, fullName := range bookings{                                  // "_" blank identifier, ignore a variable that is unused
	
						var fullNames = strings.Fields(fullName)
						firstNames = append(firstNames, fullNames[0])
					}
						fmt.Printf("The first names of bookings %v\n",firstNames)
						fmt.Printf("These are all our bookings %v\n",bookings)
						continue
				}else if userTickets > totalConferenceTickets || userTickets > int(remainingTickets){
						fmt.Printf("We only have %v of total tickets and %v are remaining.\n",totalConferenceTickets,remainingTickets)	
						continue
					
				}else {
						//fmt.Printf("We only have %v of total tickets and %v are remaining.\n",totalConferenceTickets,remainingTickets)
						
						if !isValidName{
							fmt.Printf("First name or last name is too short\n")
						}

						if !isValidEmail{
							fmt.Printf("email address doesnt contain @ sign\n")
						}

						if !isValidUserTickets{
							fmt.Printf("Number of tickets entered are invalid\n")
						}
				}		
		} 

		if remainingTickets == 0{
			fmt.Printf("Tickets are sold out, see you next year\n")
		}
	// fmt.Printf("Whole array %v \n", bookings)
	// fmt.Printf("First element of the array %v\n",bookings[0])
	// fmt.Printf("Type of the array %T \n", bookings)
	// fmt.Printf("Length of the array %v \n", len(bookings))

	// Switch statements
	// var city = "Mysore"

	// switch city {
	// 	case "Bangalore, Tumkur, Mysore":
	// 		// code for bangalore 
	// 	case "Mumbai, Pune":
	// 		// code for mumbai
	// 	case "Delhi":
	// 		// code for delhi
	// 	case "Kolkata":
	// 		// code for kolkata
	// 	case "Chennai, Coimbatore":
	// 		// code for chennai
	// 	default:
	// 		// if city doesnt match to the above
	// }
}

func greetUsers(confName string, totTickets int, remTickets uint) {
	fmt.Print("Welcome to ", confName, " booking application \n")
	fmt.Printf("We have total of %v tickets and remaining tickets are %v \n", totTickets,remTickets)
	fmt.Println("Get your tickets here to attend")
	
}
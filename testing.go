// package main

// import "fmt"

// func main(){

// 	// // Maps - only stores same datatypes
// 	// var userData = make(map[string]string)

// 	// userData["firstName"] = "Prajwal"
// 	// userData["lastName"] = "T"
// 	// userData["email"] = "prajwal@yahoo.com"

// 	// for key, value := range userData {

// 	// 	fmt.Printf("%v : %v\n", key,value)

// 	// }

// 	// creating a struct - basically creating custom data type
// 	type UserData struct {
// 		 firstName string
// 		 lastName string
// 		 email string
// 		 numberOfTickets uint
// 	}

// 	// creating a variable of struct type
// 	var userData UserData

// 	// assigning values to struct items
// 	userData.firstName = "Prajwal"
// 	userData.lastName = "t"
// 	userData.email = "tkushal@gmail.com"
// 	userData.numberOfTickets = 10

// 	// printing the values
// 	var message = fmt.Sprintf("Hi %v!!,\n\nThank you for booking %v tickets.\nWe will also email you the tickets to email address %v \n", userData.firstName, userData.numberOfTickets, userData.email )
// 	fmt.Printf(message)
// }

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main()  {
	// Prints 3 lines
	// 1st line = 10 stars
	// 2nd line = Welcome!
	// 3rd line = 10 stars
	fmt.Println(strings.Repeat("*",10) +"\n" + "Welcome!\n" + strings.Repeat("*",10)) 

	// Remove all the starts first, then remove trailing spaces from left and right.
	msg := `
	**************************
	*    BUY NOW, SAVE 10%   *
	**************************`

	fmt.Println(msg)

	afterRemovingStars := strings.ReplaceAll(msg,"*","")
	res := strings.TrimSpace(afterRemovingStars)
	fmt.Println(res)

	carsList := []string{ "Wuling Hongguang","Toyota Corolla", "Toyato Abc"}

	sort.Strings(carsList) 
	
	fmt.Println(carsList)

}


package main

import "fmt"

func main() {
	fmt.Println("This is Arrays! Not much used in Golang. Use slices!!! Next version of arrays in golang \n\n\n")

	// create an array. Dont forget to mention the size of the array
	var cityList [5]string

	var item string

	fmt.Printf("Enter 5 cities \n")

	// Get input from users and add items to array
	for i := 0; i < len(cityList); i++ {
		fmt.Printf("%v. ",i+1)
		fmt.Scan(&item)

		cityList[i] = item
		fmt.Printf("\n")
		
	}

	fmt.Printf("Array items are %v \n",cityList)


	// Create and initialize array
	var namesList = [4]string{"prajwal","kushal","ram", "rahul"}

	fmt.Println(namesList)


}

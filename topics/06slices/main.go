package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Golang slices!!! most used data type in golang")

	// create a slice
	var cityList []string 

	// add items to slice
	// cityList = append(cityList, "sira")
	// cityList = append(cityList, "tumkur")
	// cityList = append(cityList, "bangalore")
	// cityList = append(cityList, "mysore")

	// get input from users and append to slice
	var numCities int 
	var city string

	fmt.Printf("How many cities you remember? \n")
	fmt.Scan(&numCities)

	fmt.Printf("That's cool!! Plese enter those %v city names\n",numCities)

	for i := 0; i < numCities; i++ {
		fmt.Printf("%v. ",i+1)
		fmt.Scan(&city)
		cityList = append(cityList,city)
		fmt.Printf("\n")
		
	}

	fmt.Printf("Size of the slice %v\n",len(cityList))
	fmt.Printf("Contents of the slice %v\n\n",cityList)

	// slicing the slices 
	fmt.Printf("This is slicing the slices!! \n")
	//cityList = append(cityList[:2])
	fmt.Printf("cityList[:2] is %v\n",cityList[:2])   // cityList[0], cityList[1]
	fmt.Printf("cityList[1:3] is %v\n\n\n",cityList[1:3])  // cityList[1], cityList[2]


	// sorting a slice
	fmt.Printf("This is sorting the slice !! \n")
	var numList = []int{20,14,54,45,87}
	fmt.Printf("Before sort: %v\n",numList)

	// sort 
	sort.Ints(numList)
	
	fmt.Printf("After the sort: %v\n",numList)

}
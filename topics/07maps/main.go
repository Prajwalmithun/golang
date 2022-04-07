package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Welcome to golang maps!!")

	// create a map and initialize it with make() -> allocates memory 
	myMaps := make(map[string]string) 

	myMaps["Name"] = "Prajwal"
	myMaps["Email"] = "tkushal@gmail.com"
	myMaps["Age"] = "24"
	
	fmt.Printf("Here are the details of myMaps %v\n",myMaps)

	/////////////////////////////////////////////////
	// Find the frequency of letters in a string ////
	/////////////////////////////////////////////////
	word := "information"
	//fmt.Println(string(word[0]))

	// create a map 
	frequencyMap := make(map[string]int)
	
	for i := 0; i < len(word); i++ {
		frequencyMap[string(word[i])]++
	}
	//fmt.Println(len(word))

	fmt.Println(frequencyMap)

	fmt.Printf("key : frequency\n")
	fmt.Printf(strings.Repeat("=",15)+"\n")
	for key,value := range frequencyMap{
		fmt.Printf("%v : %v\n",key,value)
	}
}

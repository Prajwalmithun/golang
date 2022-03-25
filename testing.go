package main

import "fmt"

func main(){

	var userData = make(map[string]string) 

	userData["firstName"] = "Prajwal"
	userData["lastName"] = "T"
	userData["email"] = "prajwal@yahoo.com"

	for key, value := range userData {
	
		fmt.Printf("%v : %v\n", key,value)
		
	}
	
}



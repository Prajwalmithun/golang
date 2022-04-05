package main

import "fmt"

func main()  {

	var ptr *int 
	fmt.Printf("Default value stored in this pointer %v\n",ptr)
	fmt.Printf("Type of the pointer %T\n",ptr)

	someNumber := 23
	var ptr1 = &someNumber  					// pointer pointing to "someNumber"

	fmt.Printf("Address stored in pointer %v\n",ptr1)
	fmt.Printf("Value stored at address %v is %v\n",ptr1,*ptr1) 
	
}
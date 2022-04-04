package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Printf("Welcome to userinput function \n")
	
	var choice int
	fmt.Printf("1.NewReader\n2.Scanf\n")
	fmt.Scan(&choice)
	
	switch choice {

	case 1:
		fmt.Printf("This is NewReader\n")
		// Creating a reader from stdin
		reader := bufio.NewReader(os.Stdin)

		fmt.Printf("Enter your age: \n")

		// comma ok || err ok
		// if we dont care about errors then use "_"
		input, _ := reader.ReadString('\n')

		fmt.Printf("Thank you, your age is %v", input)
		fmt.Printf("Type of input %T",input)
		break

	case 2:
		// scanf -> only to take strings input.
		fmt.Printf("This is scanf\n")
		var name string
		fmt.Printf("Enter your name\n")
		fmt.Scanf("%v",&name)          // using scanf
		// fmt.Scan(&name)             // using scan

		fmt.Printf("SO you are %v \n",name)
		break

	default:
		fmt.Printf("Wrong choice\n")
		break

	}
	
}


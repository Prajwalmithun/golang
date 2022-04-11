/* 
	defer
	: it will be executed at the end of main()
	:follows LIFO

*/


package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer in golang!!")


	fmt.Println("Hello")
	fmt.Println("World")
	// output of above would be: Welcome to defer in golang!!, Hello \n World

	defer fmt.Println("Hi")
	defer fmt.Println("Nice to meet you")
	fmt.Println("How are you?")
	// Output of above would be : Welcome to defer in golang!!, Hello \n World \n How are you? \n Nice to meet you \n Hi

	myDefer()
	// output of above : Welcome to defer in golang!!, Hello, world, How are you?, 43210, Nice to meet you, Hi

	defer myDefer()
	// output : Welcome to defer in golang!!, Hello, world, How are you?, 43210, 43210, Nice to meet you, Hi

}

func myDefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
		
	}
}


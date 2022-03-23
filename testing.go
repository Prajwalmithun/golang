package main

import "fmt"

func main(){
	fmt.Printf("Sum of 1 and 2 is %v \n",add_numbers(1,2))
}

func add_numbers (a int, b int) int {
	return (a + b)
}

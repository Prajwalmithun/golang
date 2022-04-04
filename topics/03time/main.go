package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Printf("Welcome to time package \n")

	// print the present time
	presentTime := time.Now()
	fmt.Printf("%v\n",presentTime)

	// formating the above date
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))     // output: today's date time day

	fmt.Printf("%v-%v-%v\n",presentTime.Year(),presentTime.Month(),presentTime.Day())  // Output: 2022-April-4

	time.Sleep(10 * time.Second)

}

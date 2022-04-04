package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Printf("Welcome to our Shopping App Survery\n")
	fmt.Printf("Please rate your experience between 1 and 10\n")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Printf("Thank you for rating %v", input)

	// do some operation on taken input
	// Type conversion
	// strings -> float 64
	input1, err := strconv.ParseFloat(strings.TrimSpace(input),64)

	// strings -> int 64, base=decimal
	//input1, err := strconv.ParseInt(strings.TrimSpace(input),10,64)

	// strings -> int 0, base=10
	//input1, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Printf("Converted %v",input1 + 1)
	}

}

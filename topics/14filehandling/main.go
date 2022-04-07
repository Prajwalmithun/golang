package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("This is file management with Golang!!")

	fileName := "./testing.txt"
	contents := "Hello, gophers!!"

	// Create a file
	file, err := os.Create(fileName)
	checkNilError(err)

	// Write contents to that file
	length, err := io.WriteString(file, contents)
	checkNilError(err)

	fmt.Printf("Length of file is :%v\n", length)

	// Read contents from the file
	databytes, err := ioutil.ReadFile(fileName)
	checkNilError(err)

	fmt.Printf("Contents of the file is: %v\n", databytes)
	fmt.Printf("Contents of the file is: %v\n", string(databytes)) //type conversion bytes -> strings

	// Check the permissions, metadata related to that file
	info, err := os.Stat(fileName)
	checkNilError(err)

	fmt.Printf("Metadata of file: %v\n", info)

	// Close the file
	defer file.Close()

	// Delete a file
	er := os.Remove(fileName)
	checkNilError(er)

}

// Function to check if any errors occurs
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}

}

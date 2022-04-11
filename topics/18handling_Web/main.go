package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println(strings.Repeat("##",20))
	fmt.Println("Welcome to handling Web Req/Res from FrontEnd (here its from FastAPI)")
	fmt.Println(strings.Repeat("##",20))

	handleGetRequest()
	makePostJsonRequest()
	makePostFormRequest()

}

// Function to handle all sorts of err
func checkNilError(err error)  {
	if err != nil{
		panic(err)
	}
	
}

// Handling GET request
func handleGetRequest()  {
	const url = "http://localhost:8000/get"

	// make a get request
	response, err := http.Get(url)

	// handle the err
	checkNilError(err)

	// Close the connection
	defer response.Body.Close()

	databytes,err := ioutil.ReadAll(response.Body)
	checkNilError(err)


	// check the return type and return content
	fmt.Printf("The type of the response is %T: \n\n",response)     //*http.Response
	fmt.Printf("Content of the response: \n%v \n",string(databytes))

}


func makePostJsonRequest()  {
	const url = "http://localhost:8000/createItems/"

	// fake json payload 
	payload := strings.NewReader(`
	{
		"name": "Rice",
		"description": "Rice Sona Masoori",
		"price" : 100.76,
		"tax" : 3.4
	}
	`)

	response, err := http.Post(url,"application/json",payload)

	checkNilError(err)

	defer response.Body.Close()

	databytes,err := ioutil.ReadAll(response.Body)
	checkNilError(err)

	fmt.Printf("Content of the response: \n%v \n",string(databytes))

}

func makePostFormRequest()  {
	url := ""

	data := url.Values{}

	data.Add("firstname","Prajwal")
	data.Add("lastname","Thippeswamy")
	data.Add("email","prajwal@gmail.com")

	response, err := http.PostForm(url,data)

	checkNilError(err)

	defer response.Body.Close()

	contents, _ := ioutil.ReadAll(response.Body)


	fmt.Println(string(contents))
	
}
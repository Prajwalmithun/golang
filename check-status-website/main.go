package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main(){

	website_list := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.fb.com",
		"https://www.amazon.com",
		"https://www.netflix.com",
		"https://www.apple.com",
	}

	for _, web := range website_list {
		wg.Add(1)
		go getStatusCode(web)
	}
	
	wg.Wait()
}

func getStatusCode(endpoint string){
	res, err:= http.Get(endpoint)

	if err != nil {
		fmt.Printf("Cannot connect to the endpoint\n")
	}

	// fmt.Println(res.Request)
	// fmt.Println()
	fmt.Printf("%d OK | %s\n",res.StatusCode,endpoint)

	defer wg.Done()
}

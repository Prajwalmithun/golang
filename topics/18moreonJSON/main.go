package main

import (
	"encoding/json"
	"fmt"
)

type item struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json: price`
	Tax         float64 `json: tax`
	// password string `JSON:"-"`
}

func main() {
	fmt.Println("Encode JSON, Decode JSON ")

	//encodeJson()
	decodeJson()

}

func checkNillError(err error) {
	if err != nil {
		panic(nil)
	}
}

// Create JSON
func encodeJson() {

	listOfItems := []item{
		{0, "sugar", "makes your life sweet!!", 40, 2},
		{1, "chilli", "makes your life hot!!", 140, 20},
		{2, "lemon", "makes your life cool!!", 20, 0.2},
	}

	// finalJson, err := json.Marshal(listOfItems)      # <<<-- Output is not humanble readable
	finalJson, err := json.MarshalIndent(listOfItems, "", "\t")

	checkNillError(err)

	fmt.Printf("%s\n", finalJson)

}

// Decode JSON
func decodeJson(){

	jsonFromWeb := []byte(`
	{
		"id": 0,
		"name": "sugar",
		"description": "makes you sweet!",
		"price": 40.0,
		"tax": 2
	}
	`)

	var retailItem item

	validJson := json.Valid(jsonFromWeb)

	if validJson{
		json.Unmarshal(jsonFromWeb, &retailItem)
		fmt.Printf("%#v\n",retailItem)
	}else{
		fmt.Println("Invalid JSON")
	}


	// for computation of JSON we received, we can use maps

	var myItem map[string]interface{}

	json.Unmarshal(jsonFromWeb, &myItem)
	fmt.Println(myItem)


	// loop over the items in [json]maps

	for k,v := range myItem{
		fmt.Printf("Key is %v, and key is %v\n",k,v)
	}
}
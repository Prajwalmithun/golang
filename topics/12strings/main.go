package main

import (
	"fmt"
	"sort"
	"strings"
)

func main()  {
	// Prints 3 lines
	// 1st line = 10 stars
	// 2nd line = Welcome!
	// 3rd line = 10 stars
	fmt.Println(strings.Repeat("*",10) +"\n" + "Welcome!\n" + strings.Repeat("*",10)) 

	// Remove all the starts first, then remove trailing spaces from left and right.
	msg := `
	**************************
	*    BUY NOW, SAVE 10%   *
	**************************`

	fmt.Println(msg)

	afterRemovingStars := strings.ReplaceAll(msg,"*","")
	res := strings.TrimSpace(afterRemovingStars)
	fmt.Println(res)

	
}
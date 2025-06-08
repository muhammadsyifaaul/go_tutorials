package main

import (
	"errors"
	"fmt"
)

func main(){
	var myString string = "Hello"
	fmt.Println(myString)

	var myString2 string = "World"
	fmt.Println(myString2)

	var myString3 string = myString + myString2
	fmt.Println(myString3)

	var myString4 string = "Hello"
	fmt.Println(myString4)

	var myString5 string = "World"
	fmt.Println(myString5)

	var myString6 string = myString4 + myString5
	fmt.Println(myString6)

	var myString7 string = "Hello"
	fmt.Println(myString7)

	var myString8 string = "World"
	fmt.Println(myString8)

	myResult, err := calculate(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(myResult)
	}
	
}

func calculate(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
// This code demonstrates variable declaration, string concatenation, and error handling in Go.

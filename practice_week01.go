package main

import "fmt"

func test05() {
	myString := [5]string{"I","am","stupid","and","weak"}
	fmt.Printf("myString %+v\n", myString)
	for index, value := range myString {
		if value == "stupid" {
			myString[index] = "smart"
		}
		if value == "weak" {
			myString[index] = "strong"
		}

	}
	fmt.Printf("myString %+v\n", myString)
}

func main() {
	test05()
}

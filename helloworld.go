package main

import (
	"flag"
	"fmt"
	"os"
)

func test01() {
	name := flag.String("name", "world", "specify the name you want to say hi")
	flag.Parse()
	fmt.Println("OS args is:", os.Args)
	fmt.Println("input parameter is:", *name)
	fullString := fmt.Sprintf("Hello %s form Go/n", *name)
	fmt.Println(fullString)
}

func test02() {
	fullString := "hello world"
	fmt.Println(fullString)
	for i, c := range fullString {
		fmt.Println(i, string(c))
	}
}

func test03() {
	c := []int{6, 7, 8} // creates and array and returns a slice reference
	fmt.Println(c)
}

func test04() {
	mySlice := []int{10, 20, 30, 40, 50}
	for _, value := range mySlice {
		value *= 2
	}
	fmt.Printf("mySlice %+v\n", mySlice)
	for index := range mySlice {
		mySlice[index] *= 2
	}
	fmt.Printf("mySlice %+v\n", mySlice)
}


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
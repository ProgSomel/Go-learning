package main

import "fmt"

func printWelcomeMessage(){
	fmt.Println("Welcome to the application")
}

func getUserName() string {
	var name string

	fmt.Println("Enter your name-")
	fmt.Scanln(&name)

	return name
}

func getTwoNumbers() (int, int){
	var num1 int
	var num2 int

	fmt.Println("Enter First Number")
	fmt.Scanln(&num1)

	fmt.Println("Enter Second Number")
	fmt.Scanln(&num2)

	return num1, num2
}

func add(num1 int, num2 int) int {
	sum:= num1 + num2
	return sum
}

func display(name string, sum int){
	// Display Result 
	fmt.Println("Hello, ", name)
	fmt.Println("Summation = ", sum)
}

func goodBye(){
	// GoodBye message 
	fmt.Println("Thank you for using the application")
	fmt.Println("Good Bye")
}

func main(){
	printWelcomeMessage()
	name:= getUserName()
	num1, num2:=getTwoNumbers()
	sum:=add(num1, num2)
	display(name, sum)
	goodBye()
}
	

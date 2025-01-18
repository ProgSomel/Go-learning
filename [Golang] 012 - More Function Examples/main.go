package main

import "fmt"

func add(num1 int, num2 int) int {
	sum:=num1 + num2

	return sum
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum:=num1 + num2
	mul:= num1 * num2

	return sum, mul
}

func printSomething(){
	fmt.Println("Education must be Free")
}

func sayHello(name string){
	fmt.Println("Welcome to the GoLang course, ", name)
}

func main(){
	printSomething()
	sayHello("Somel")
}
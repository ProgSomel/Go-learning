package main

import "fmt"

func main(){
	// Function expression or Assign Function in variable 
	add(2, 3)

	add:= func(a int, b int){
		c:= a + b
		fmt.Println(c)
	}

}

func init(){
	fmt.Println("I will be called first")
}
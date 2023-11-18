package main

import (
	"fmt"
)

func main() {
	calculator(10, 0, "/")
}

func calculator(num1, num2 int, operator string) {
	switch operator {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		if num2 == 0 && operator == "/" {
		fmt.Println("На ноль делить нельзя, идиотина")
		} else {
			fmt.Println(num1 / num2)
		}
	case "%":
		fmt.Println(num1 % num2)
	default:
		fmt.Println("Ты даун?")
	}

	
}

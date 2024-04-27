package main

import (
	"fmt"

	"project.go/utility"
)

func main() {
	var name string
	fmt.Println("Qual seu nome?")
	fmt.Scan(&name)
	fmt.Println("Prazer em te conhecer", name)
	fmt.Println("Qual palavra devo escrever o reverso?")
	fmt.Scan(&name)
	fmt.Println(utility.Reverse(name))
}

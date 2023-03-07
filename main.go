package main

import "fmt"

func main() {
	a := 10
	b := 20

	if a > b {
		fmt.Println("A maior que B")
	} else if b > a {
		fmt.Println("B maior que A")
	} else {
		fmt.Println("São iguais")
	}
	// print("Hello world!")
}

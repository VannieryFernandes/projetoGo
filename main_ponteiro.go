package main

import "fmt"

func main() {
	a := 10
	b := &a //Referencia da memorida
	// b = 20
	//referenciando a memoria com "&"", depois pego o valor indo até a memoria com asterisco
	fmt.Println(a, *b)
	a = 20
	// *b = 50
	fmt.Println(a, *b)

}

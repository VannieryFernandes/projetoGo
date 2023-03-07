package main

import "fmt"

type Order struct {
	ID         string
	price      float64
	quantidade int
}

func (o Order) getTotal() float64 {
	return o.price * float64(o.quantidade)
}

func main() {
	order := Order{
		ID:         "123",
		price:      10.5,
		quantidade: 8,
	}

	fmt.Println(order.ID, order.price, order.quantidade)
	fmt.Println(order.getTotal())

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

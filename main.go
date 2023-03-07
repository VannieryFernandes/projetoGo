package main

import "fmt"

type Order struct {
	ID       string
	Price    float64
	Quantity int
}

func (o Order) GetTotal() float64 {
	return o.Price * float64(o.Quantity)
}

func (o *Order) SetPrice(price float64) { // asterisco é ponteiro, deixando a função habilitada para setar valor da origem
	o.Price = price
	fmt.Println("Price dentro do SetPrice", o.Price)
}

func main() {
	order := Order{
		ID:       "123",
		Price:    10,
		Quantity: 8,
	}

	//TESTES COM FUNÇÃO E PRINTS
	fmt.Println(order.ID, order.Price, order.Quantity)
	fmt.Println(order.GetTotal())

	//Testes com setPrice
	order.SetPrice(20)
	fmt.Println("Preço Total", order.GetTotal())

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

package main

import "fmt"

func main() {

	var SumPac int
	var SumTrade int

	registerTrade()
	fmt.Printf("Total invested: %v", SumPac+SumTrade)
}

func registerTrade() {
	//  qui  dobbiamo fare in modo che un valore investito viene registrato e assegnato ad un investimento

	var asset string
	var q int
	var pEntry int
	var pExit int

	fmt.Println("What asset trade do you want to register?")
	fmt.Scan(&asset)
	fmt.Println("How much did you invest and at what price?")
	fmt.Scan(&q, &pEntry)
	fmt.Printf("il valore attuale di questa trade è %v, a fronte di un investimento iniziale %v, per un rendimento di %v", q, q*pEntry, q*pExit-q*pEntry)

}

//make a func to register pac

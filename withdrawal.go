package main

type Withdrawal struct {
	ID    string
	Itens []Item
}

type Item struct {
	Product string
	Amount  int
}

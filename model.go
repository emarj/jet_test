package main

type Entity struct {
	ID   int    `json:"id" sql:"primary_key"`
	Name string `json:"name"`
}

type Account struct {
	ID    int    `json:"id" sql:"primary_key"`
	Name  string `json:"name"`
	Owner Entity `json:"owner" alias:"owner"`
}
type Operation struct {
	ID     int     `json:"id" sql:"primary_key"`
	From   Account `json:"from" alias:"from"`
	To     Account `json:"to" alias:"to"`
	Amount int     `json:"amount"`
}

package main

import (
	"math/rand"
	"time"
)


type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`

}

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func newAccount(firstname, lastname string) *Account {
	return &Account{
		FirstName: firstname,
		LastName:  lastname,
		Number:    int64(rand.Intn(1000000)),
		CreatedAt: time.Now().UTC(),
	}
}

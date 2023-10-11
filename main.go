package main

import (
	"flag"
	"log"
)

func seedAccount(store Storage,fname, lname, pw string) *Account {
	acc, err := newAccount(fname, lname, pw)
	if err!= nil {
        log.Fatal(err)
    }

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "Anthone", "GG", "cacador123188")
}

func main() {
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostGressStore()

	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		seedAccounts(store)
	}

	server := NewApiServer(":3000", store)
	server.run()
}
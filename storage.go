package main

import (
	"database/sql"
	_"github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(id int) error
	UpdateAccount(*Account) error
	GetAccountById(id int) (*Account, error)
}


type PostGressStore struct {
	db *sql.DB
}

func NewPostGressStore()(*PostGressStore, error) {
	connStr := "user=postgres dbname=postgres password=api-golang sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostGressStore{
		db: db,
	}, nil 

}

func (s *PostGressStore) Init() error {
	return s.CreateAccountTable()
}

func (s *PostGressStore) CreateAccountTable() error {
	query := `create table if not exists account (
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		number serial,
		balance serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)

	return err
}






func (s *PostGressStore) CreateAccount(*Account) error {
	// resp, err := s.db.Query(`insert into account values ()`)
	return nil
}

func (s *PostGressStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostGressStore) DeleteAccount(id int) error {
	return nil
}

func (s *PostGressStore) GetAccountById(id int) (*Account, error) {
	return nil, nil
}
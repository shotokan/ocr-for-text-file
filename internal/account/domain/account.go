package domain

import (
	"errors"

	"github.com/google/uuid"
)

type Account struct {
	id string
	number string
	checksum string
}


func NewAccount(id, number string) (*Account, error ) {
	if id == "" {
		return nil, errors.New("empty account id")
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("not valid account id")
	}

	if number == "" {
		return nil, errors.New("empty account number")
	}

	return &Account{
		id: id,
		number: number,
	}, nil
}

func (a Account) ID() string {
	return a.id
}

func (a Account) Number() string {
	return a.number
}

func (a Account) CheckSum() string {
	return a.checksum
}

func (a *Account) Validate() {
	a.checksum = checkSum(a.number)
}
package common

import "math/big"

type UserDeposit struct {
	Tx     string
	UserId int
	Amount *big.Int
}

type NewTx struct {
	Id     int
	Type   int
	Status int
	Debit  int
	Credit int
	UserId int
}

package common

import "math/big"

type UserDeposit struct {
	Tx       string
	WalletId int
	Amount   *big.Int
}

type NewTx struct {
	Id       int
	Type     int
	Status   int
	Debit    int
	Credit   int
	WalletId int
}

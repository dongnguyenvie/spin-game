package transactionmodel

type TransactionDeposit struct {
	Tx      string `json:"tx" gorm:"column:tx;"`
	Type    int    `json:"type" gorm:"column:type;"`
	Credit  int64  `json:"credit" gorm:"column:credit;"`
	Debit   int    `json:"debit" gorm:"column:debit;"`
	Status  int    `json:"status" gorm:"column:status"`
	Package int    `json:"package" gorm:"column:package"`

	WalletId int `json:"walletId" gorm:"column:wallet_id;"`
}

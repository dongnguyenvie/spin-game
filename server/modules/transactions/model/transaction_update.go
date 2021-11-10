package transactionmodel

type TransactionUpdateStatus struct {
	Status int `json:"status" gorm:"column:status"`
}

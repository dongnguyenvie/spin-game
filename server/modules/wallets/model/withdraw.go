package walletmodel

type Withdraw struct {
	Tx     string `json:"tx"`
	Amount int64  `json:"amount"`
}

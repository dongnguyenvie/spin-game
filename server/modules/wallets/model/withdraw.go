package walletmodel

type Withdraw struct {
	Tx     string `json:"tx"`
	Amount int    `json:"amount"`
}

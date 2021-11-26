package gamespinmodel

type BuyPackage struct {
	Quantity int `json:"quantity"`
}

func (model *BuyPackage) Fullfill() {
	if model.Quantity < 0 {
		model.Quantity = 1
	}
}

package usermodel

type Filter struct {
	OwnerId int   `json:"owner_id,omitempty"`
	Status  []int `json:"-"`
}

package common

import "time"

type SQLModel struct {
	Id       int        `json:"id" gorm:"column:id;"`
	FakeId   string     `json:"FakeId" gorm:"-;"`
	Status   int        `json:"status" gorm:"column:status;default:1;"`
	CreateAt *time.Time `json:"createAt,omitempty" gorm:"column:created_at;"`
	UpdateAt *time.Time `json:"updateAt,omitempty" gorm:"column:update_at;"`
}

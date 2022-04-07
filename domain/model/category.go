package model

import "time"

const (
	CategoryDelete = 100
	CategoryInUse  = 1
)

type Category struct {
	ID          uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name        string `gorm:"type:varchar(20);not null" json:"name"`
	State       int    `gorm:"type:bigint(20);not null" json:"state"`
	CreatedTime time.Time
	UpdatedTime time.Time
}

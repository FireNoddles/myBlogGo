package model

import (
	"github.com/jinzhu/gorm"
)

const (
	UserDelete = 100
)

type Users struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	State    int    `gorm:"type:int;DEFAULT:1" json:"state" validate:"required" label:"状态"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

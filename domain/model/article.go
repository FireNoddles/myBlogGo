package model

import (
	"time"
)

const (
	ArticleDelete = 100
	ArticleInUse  = 1
)

type Article struct {
	ID          uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name        string `gorm:"type:varchar(20);not null" json:"name"`
	CreatedTime time.Time
	UpdatedTime time.Time
	Cid         int    `gorm:"type:int;not null" json:"cid"`
	Desc        string `gorm:"type:varchar(200)" json:"desc"`
	Content     string `gorm:"type:longtext" json:"content"`
	Img         string `gorm:"type:varchar(200)" json:"img"`
	CmtCount    int    `gorm:"type:int;not null;default:0" json:"cmt_count"`
	ReadCount   int    `gorm:"type:int;not null;default:0" json:"read_count"`
	State       int    `gorm:"type:int;not null;default:0" json:"read_count"`
}

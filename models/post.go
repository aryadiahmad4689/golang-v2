package models

import (
	"github.com/jinzhu/gorm"
)

// Posts ..
type Posts struct {
	gorm.Model
	Title string `gorm:"type:varchar(50)"`
	Desc  string `gorm:"type:text"`
}

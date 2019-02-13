package entity

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Todo struct {
	gorm.Model
	Name        string `gorm:"type:varchar(50)"`
	Description string `gorm:"type:varchar(200)"`
}

package models

import (
	"github.com/jinzhu/gorm"
)

type Url struct {
	gorm.Model
	Original string
	Count    int
}

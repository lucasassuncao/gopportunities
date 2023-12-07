package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Link     string
	Location string
	Remote   bool
	Salary   int64
}

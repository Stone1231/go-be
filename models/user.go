package models

import (
	"time"
)

type User struct {
	//gorm.Model
	ID          uint      `gorm:"primary_key;AUTO_INCREMENT" json:"id" uri:"id"`
	Name        string    `json:"name"`
	Hight       int32     `gorm:"hight" json:"hight"`
	Birthday    time.Time `gorm:"birthday" json:"-"`
	BirthdayStr string    `gorm:"-" json:"birthday"`
	Photo       string    `gorm:"photo" json:"photo"`
	Dept        Dept      `gorm:"foreignkey:DeptID;" json:"-"`
	// DeptID      uint      `gorm:"dept_id" json:"dept,string,omitempty"`
	DeptID  uint   `gorm:"dept_id" json:"dept"`
	Projs   []Proj `gorm:"many2many:user_proj;" json:"-"`
	ProjIDs []uint `gorm:"-" json:"projs"`
}

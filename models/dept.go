package models

type Dept struct {
	ID    uint   `gorm:"primary_key" json:"id" uri:"id"`
	Name  string `json:"name"`
	Users []User `gorm:"foreignkey:DeptID" json:"users"`
}

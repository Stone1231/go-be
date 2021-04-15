package models

type Proj struct {
	ID    uint   `gorm:"primary_key" json:"id" uri:"id"`
	Name  string `json:"name"`
	Users []User `gorm:"many2many:user_proj;" json:"users"`
}

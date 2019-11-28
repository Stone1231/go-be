package models

type AdminUser struct {
	UserName string `form:"username" json:"username"`
	Role     string `form:"role" json:"role"`
}

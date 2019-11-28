package models

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"pwd" json:"pwd" binding:"required"`
}

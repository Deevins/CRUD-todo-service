package entity

type User struct {
	Id               int    `json:"-"`
	Name             string `json:"name" binding:"required"`
	Username         string `json:"username" binding:"required"`
	Password         string `json:"password" binding:"required"`
	RepeatedPassword string `json:"repeatedPassword" binding :"required"`
}

package model

type User struct {
	Id       int    `gorm:"type:int;primary_key"`
	Username string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"uniqueIndex;type:varchar(255);not null"`
	Password string `gorm:"type:varchar(255);not null"`
}

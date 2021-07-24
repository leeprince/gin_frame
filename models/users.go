package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Id int `gorm:"primaryKey"`
	Username string
	Password string
	Age int
}
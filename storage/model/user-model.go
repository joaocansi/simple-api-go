package storage

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string
	Email     string
	Password  string
	AvatarUrl string
}

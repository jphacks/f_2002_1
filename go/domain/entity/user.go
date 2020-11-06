package entity

import "time"

type User struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Uid       string
	Name      string
}

type Users []User

type Id int

package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name         string
	Email        string
	HashPassword string
}

type Websites struct {
	gorm.Model
	Id   int64
	Web  string
	Name string
}

type SOS_User struct {
	gorm.Model
	Web_id   int64
	Email    string
	Websites Websites `gorm:"foreignKey":Id`
}

type Status_code struct {
	gorm.Model
	Web_id   int64
	Code     string
	Websites Websites `gorm:"foreignKey":Id`
}
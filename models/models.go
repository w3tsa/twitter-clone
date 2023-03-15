package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name string				`json:"name" gorm:"text;not null;default:null"`
	Email string			`json:"email" gorm:"uniqueIndex;not null;default:null"`
	Password string 		`json:"password" gorm:"text;not null;default:null"`
	Bio string 				`json:"bio" gorm:"text;not null;default:null"`
	Location string 		`json:"location" gorm:"text;not null;default:null"`
	AvatarURL string 		`json:"avatarUrl" gorm:"text;not null;default:null"`
	CreatedAt time.Time 	`json:"createdAt" gorm:"text;not null;default:null"`
	UpdatedAt time.Time 	`json:"updatedAt" gorm:"text;not null;default:null"`
	Role bool 				`json:"role" gorm:"bool;not null;default:null"`
}
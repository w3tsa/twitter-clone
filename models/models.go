package models

import (
	"time"
)

type User struct {
    ID        uint `gorm:"primaryKey;autoIncrement"`
    Username  string `gorm:"unique;not null"`
    Email     string `gorm:"unique;not null"`
    Password  string `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
    Posts    []Post
    Followers []User `gorm:"many2many:user_followers;joinForeignKey:follower_id;joinReferences:following_id"`
    Following []User `gorm:"many2many:user_following;joinForeignKey:following_id;joinReferences:follower_id"`
}

type Post struct {
    ID        uint `gorm:"primaryKey;autoIncrement"`
    Content   string `gorm:"type:text;not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
    UserID    uint
    User      User
}
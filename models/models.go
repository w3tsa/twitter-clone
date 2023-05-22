package models

import (
	"gorm.io/gorm"
)

type Post struct {
    gorm.Model
    Content     string      `json:"content"`
    UserID  uint
}

type User struct {
    gorm.Model
    ID          uint
    Username    string  `json:"username" gorm:"uniqueIndex: not null"`
    Password    string  `json:"password" gorm:"not null"`
    Posts []Post
}

// type User struct {
//     gorm.Model
//     Username   string `gorm:"uniqueIndex: not null" json:"username"`
//     Email      string `gorm:"uniqueIndex: not null" json:"email"`
//     Password   string `gotm:"not null" json:"-"`
//     Bio        string `json:"bio"`
//     Avatar     string `json:"avatar"`
//     Tweets     []Tweet `json:"tweets"`
//     Followers  []*Follow `json:"followers"`
//     Following  []*Follow `json:"following"`
// }

// type Tweet struct {
//     gorm.Model
//     UserID     uint `json:"-"`
//     User       User `json:"user"`
//     Text       string `json:"text"`
//     Likes      []*Like `json:"likes"`
//     Retweets   []*Retweet `json:"retweets"`
//     Mentions   []Mention `json:"mentions"`
// }

// type Like struct {
//     gorm.Model
//     UserID     uint `json:"-"`
//     User       User `json:"user"`
//     TweetID    uint `json:"-"`
//     Tweet      Tweet `json:"tweet"`
// }

// type Retweet struct {
//     gorm.Model
//     UserID     uint `json:"-"`
//     User       User `json:"user"`
//     TweetID    uint `json:"-"`
//     Tweet      Tweet `json:"tweet"`
// }

// type Follow struct {
//     gorm.Model
//     FollowerID uint `json:"-"`
//     Follower   User `json:"follower"`
//     FolloweeID uint `json:"-"`
//     Followee   User `json:"followee"`
// }

// type Mention struct {
//     gorm.Model
//     UserID     uint `json:"-"`
//     User       User `json:"user"`
//     TweetID    uint `json:"-"`
//     Tweet      Tweet `json:"tweet"`
// }

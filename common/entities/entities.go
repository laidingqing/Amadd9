package entities

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// ContactInfo model
type ContactInfo struct {
	Phone    string `bson:"phone"  json:"phone,omitempty"`
	Mobile   string `bson:"mobile"  json:"mobile,omitempty"`
	Email    string `bson:"email"  json:"email,omitempty"`
	Location string `bson:"location"  json:"location,omitempty"`
}

// UserPublic model
type UserPublic struct {
	Name            string      `bson:"name" json:"name"`
	Title           string      `bson:"title" json:"title,omitempty"`
	Contact         ContactInfo `bson:"contactInfo" json:"contactInfo"`
	Avatar          string      `bson:"avatar" json:"avatar"`
	AvatarThumbnail string      `bson:"avatarThumbnail" json:"avatarThumbnail"`
}

// User Model
type User struct {
	ID             bson.ObjectId `bson:"_id" json:"id"`
	UserName       string        `bson:"name" json:"name"`
	Password       string        `bson:"password" json:"password,omitempty"`
	Slat           string        `bson:"slat" json:"-"`
	Roles          []string      `bson:"roles" json:"roles"`
	Type           string        `bson:"type" json:"type"`
	CreatedAt      time.Time     `bson:"createdAt" json:"createdAt,omitempty"`
	ModifiedAt     time.Time     `bson:"modifiedAt" json:"modifiedAt,omitempty"`
	PassResetToken ActionToken   `bson:"passwordReset" json:"passwordReset,omitempty"`
	Public         UserPublic    `bson:"userPublic" json:"userPublic,omitempty"`
}

// ActionToken model
type ActionToken struct {
	Token   string    `bson:"token" json:"token"`
	Expires time.Time `bson:"expires" json:"expires"`
}

type CurrentUserInfo struct {
	Roles []string
	User  *User
}

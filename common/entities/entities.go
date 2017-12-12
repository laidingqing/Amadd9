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

//WikiRecord wiki record model.
type WikiRecord struct {
	ID          bson.ObjectId `bson:"_id" json:"id,omitempty"`
	Slug        string        `bson:"slug" json:"slug,omitempty"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description,omitempty"`
	CreatedAt   time.Time     `bson:"createdAt" json:"createdAt"`
	ModifiedAt  time.Time     `bson:"modifiedAt" json:"modifiedAt"`
	HomePageID  string        `bson:"homePageId" json:"homePageId,omitempty"`
	AllowGuest  bool          `bson:"allowGuest" json:"allowGuest"`
	Type        string        `bson:"type" json:"type"`
}

type ViewResponse struct {
	TotalRows int `json:"total_rows"`
	Offset    int `json:"offset"`
}

//Artist 艺术家
type Artist struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `bson:"name" json:"name"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
}

//Revision tab版本
type Revision struct {
	RevisionID string    `bson:"_id" json:"id"`
	TabsID     string    `bson:"tabsID" json:"tabsID"`
	UserID     string    `bson:"userID" json:"userID"`
	UploadedAt time.Time `bson:"uploadedAt" json:"uploadedAt"`
	FileName   string    `bson:"fileName" json:"fileName"`
	Summary    string    `bson:"summary" json:"summary"`
}

//Tracks backing tracks
type Tracks struct {
	TrackID    string    `bson:"_id" json:"id"`
	TabsID     string    `bson:"tabsID" json:"tabsID"`
	UserID     string    `bson:"userID" json:"userID"`
	UploadedAt time.Time `bson:"uploadedAt" json:"uploadedAt"`
	FileName   string    `bson:"fileName" json:"fileName"`
	Summary    string    `bson:"summary" json:"summary"`
}

//Annotation tabs annotation.
type Annotation struct {
	AnnotationID string `bson:"_id" json:"id"`
	TabsID       string `bson:"tabsID" json:"tabsID"`
	UserID       string `bson:"userID" json:"userID"`
	//TODO other fields
}

//Tabs gtp tab libs
type Tabs struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Name       string        `bson:"name" json:"name"`
	Slug       string        `bson:"slug" json:"slug"`
	ArtistID   string        `bson:"artistID" json:"artistID"`
	CreatedAt  time.Time     `bson:"createdAt" json:"createdAt"`
	ModifiedAt time.Time     `bson:"modifiedAt" json:"modifiedAt"`
	UseRevID   string        `bson:"useRevID" json:"useRevID"`
}

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

//ArtistRecord 艺术家
type ArtistRecord struct {
	ID        bson.ObjectId `bson:"_id" json:"id,omitempty"`
	Name      string        `bson:"name" json:"name,omitempty"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt,omitempty"`
}

//RevisionRecord tab版本
type RevisionRecord struct {
	RevisionID string    `bson:"_id" json:"id"`
	TabsID     string    `bson:"tabsID" json:"tabsID"`
	UserID     string    `bson:"userID" json:"userID"`
	UploadedAt time.Time `bson:"uploadedAt" json:"uploadedAt"`
	FileName   string    `bson:"fileName" json:"fileName"`
	Summary    string    `bson:"summary" json:"summary"`
}

//TrackRecord backing tracks
type TrackRecord struct {
	TrackID    string    `bson:"_id" json:"id"`
	TabsID     string    `bson:"tabsID" json:"tabsID"`
	UserID     string    `bson:"userID" json:"userID"`
	UploadedAt time.Time `bson:"uploadedAt" json:"uploadedAt"`
	FileName   string    `bson:"fileName" json:"fileName"`
	Summary    string    `bson:"summary" json:"summary"`
}

//AnnotationRecord tabs annotation.
type AnnotationRecord struct {
	AnnotationID string `bson:"_id" json:"id"`
	TabsID       string `bson:"tabsID" json:"tabsID"`
	UserID       string `bson:"userID" json:"userID"`
	//TODO other fields
}

//TabRecord gtp tab libs
type TabRecord struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Name       string        `bson:"name" json:"name,omitempty"`
	Slug       string        `bson:"slug" json:"slug,omitempty"`
	ArtistID   string        `bson:"artistID" json:"artistID,omitempty"`
	Artist     ArtistRecord  `bson:"-" json:"artist,omitempty"`
	CreatedAt  time.Time     `bson:"createdAt" json:"createdAt,omitempty"`
	ModifiedAt time.Time     `bson:"modifiedAt" json:"modifiedAt,omitempty"`
	UseRevID   string        `bson:"useRevID" json:"useRevID,omitempty"`
}

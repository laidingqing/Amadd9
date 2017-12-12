package user_service

import (
	"log"

	. "github.com/laidingqing/amadd9/common/auth"
	"github.com/laidingqing/amadd9/common/config"
	. "github.com/laidingqing/amadd9/common/database"
	. "github.com/laidingqing/amadd9/common/entities"
	"github.com/laidingqing/amadd9/common/services"
	"github.com/laidingqing/amadd9/common/util"
	couchdb "github.com/rhinoman/couchdb-go"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	userDbCollection = "users"
)

type Registration struct {
	NewUser User `json:"user"`
}

type UserManager struct{}

// SetUp
// Do't use , tests only
func (um *UserManager) SetUp(registration *Registration) (string, error) {
	return "", nil
}

//Register a new user
func (um *UserManager) Register(newUser *User) (string, error) {
	adminUser := services.GetAdminUser()
	return um.Create(newUser, adminUser)
}

//Create a normal user
func (um *UserManager) Create(newUser *User, curUser *CurrentUserInfo) (string, error) {
	theUser := curUser.User
	if !util.HasRole(theUser.Roles, "admin") {
		return "", NotAdminError()
	}
	query := func(c *mgo.Collection) error {
		newUser.ID = bson.NewObjectId()
		newUser.Password = CalculatePassHash(newUser.Password, newUser.Slat)
		return c.Insert(newUser)
	}

	err := ExecuteQuery(userDbCollection, query)
	log.Printf("Creating new user account for: %v", newUser.UserName)
	return newUser.ID.Hex(), err
}

//findByUsername find a user
func (um *UserManager) checkUserByUsername(theUser *User) (string, error) {
	err := um.validateUser(theUser)
	if err != nil {
		return "", err
	}
	var isUser User
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"name": theUser.UserName}).One(&isUser)
	}
	err = ExecuteQuery(userDbCollection, query)
	if err != nil {
		return "", err
	}

	log.Printf("password: %v, %v, %v", isUser.Password, theUser.Password, CalculatePassHash(theUser.Password, isUser.Slat))

	if isUser.Password != CalculatePassHash(theUser.Password, isUser.Slat) {
		return "", &couchdb.Error{
			StatusCode: 401,
			Reason:     "Username or Password Incorrect",
		}
	}

	return isUser.ID.Hex(), nil
}

// validateUser func
func (um *UserManager) validateUser(user *User) error {
	var err error
	if user.UserName == "" || len(user.UserName) < 3 || len(user.UserName) > 80 {
		err = &couchdb.Error{
			StatusCode: 400,
			Reason:     "Username invalid",
		}
	}
	return err
}

//Validate Password
func (um *UserManager) validatePassword(password string) error {
	if len(password) < config.Auth.MinPasswordLength {
		return &couchdb.Error{
			StatusCode: 400,
			Reason:     "Password too short",
		}
	} else {
		return nil
	}
}

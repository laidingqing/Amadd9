package auth

import (
	"fmt"
	"net/http"
	"time"
)

type Session struct {
	ID        string    `json:"id"`
	User      string    `json:"user"`
	Roles     []string  `json:"roles"`
	AuthType  string    `json:"authType"`
	ExpiresAt time.Time `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserLoginCredentials struct {
	Username string `json:"name"`
	Password string `json:"password"`
	AuthType string `json:"auth_type"`
}

type Authenticator interface {
	CreateSession(string, string) (*Session, error)
	DestroySession(string) error
}

type AuthError struct {
	ErrorCode int
	Reason    string
}

func UnauthenticatedError() error {
	return &AuthError{
		ErrorCode: 401,
		Reason:    "Invalid username or password",
	}
}

func (err *AuthError) Error() string {
	return fmt.Sprintf("[Error]:%v: %v", err.ErrorCode, err.Reason)
}

func GetBasicAuth(req *http.Request) couchdb.Auth {
	if req.Header.Get("Authorization") != "" {
		return &couchdb.PassThroughAuth{
			AuthHeader: req.Header.Get("Authorization"),
		}
	} else {
		return nil
	}
}

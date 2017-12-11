package auth

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

var (
	//SecretKey JWT secret key
	SecretKey = "welcome to amadd9"
)

type JwtSession struct {
	ID    string `json:"id"`
	User  string `json:"user"`
	Token string `json:"token"`
}

type Session struct {
	ID    string   `json:"id"`
	User  string   `json:"user"`
	Roles []string `json:"roles"`
	Token string   `json:"token"`
}

type AuthError struct {
	ErrorCode int
	Reason    string
}

// UnauthenticatedError .auth error
func UnauthenticatedError() error {
	return &AuthError{
		ErrorCode: 401,
		Reason:    "Invalid username or password",
	}
}

// InvalidTokenError .auth error
func InvalidTokenError() error {
	return &AuthError{
		ErrorCode: 401,
		Reason:    "Token is invalided",
	}
}

func (err *AuthError) Error() string {
	return fmt.Sprintf("[Error]:%v: %v", err.ErrorCode, err.Reason)
}

// CreateJWT generat a jwt token
func CreateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	return token.SignedString([]byte(SecretKey))
}

// ValidateTokenMiddleware ...
func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

	if err == nil {
		if token.Valid {
			next(w, r)
		} else {
			log.Println("Token is not valid")
			w.WriteHeader(http.StatusUnauthorized)
		}
	} else {
		log.Println("Unauthorized access to this resource")
		w.WriteHeader(http.StatusUnauthorized)
	}

}

// ValidateToken Simple validate jwt token
func ValidateToken(r *http.Request) bool {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

	if err == nil {
		if token.Valid {
			return true
		} else {
			log.Println("Token is not valid")
			return false
		}
	} else {
		log.Println("Unauthorized access to this resource")
		return false
	}

}

// CalculatePassHash .
func CalculatePassHash(pass, salt string) string {
	h := sha1.New()
	io.WriteString(h, salt)
	io.WriteString(h, pass)
	return fmt.Sprintf("%x", h.Sum(nil))
}

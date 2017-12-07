package err

import (
	"fmt"
)

type Error struct {
	StatusCode int
	URL        string
	Method     string
	ErrorCode  string
	Reason     string
}

func NotAdminError() error {
	return &Error{
		StatusCode: 403,
		Reason:     "Not an admin",
	}
}

//stringify the error
func (err *Error) Error() string {
	return fmt.Sprintf("[Error]:%v: %v %v - %v %v",
		err.StatusCode, err.Method, err.URL, err.ErrorCode, err.Reason)
}

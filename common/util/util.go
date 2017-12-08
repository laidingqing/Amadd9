package util

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

//Returns the directory the executable is running from
func GetExecDirectory() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}

//Returns the "default" config file name and location
func DefaultConfigLocation() (string, error) {
	dir, err := GetExecDirectory()
	if err != nil {
		return "", err
	} else {
		return path.Join(dir, "config.ini"), nil
	}
}

//Just a package of utility functions that don't really belong
//anywhere else

func ContainsString(findStr string, strArr []string) bool {
	for _, str := range strArr {
		if str == findStr {
			return true
		}
	}
	return false
}

func HasRole(roles []string, role string) bool {
	return ContainsString(role, roles)
}

// Check if this user is an admin of anything.
// Either a site admin, or a wiki admin returns true
func IsAnyAdmin(roles []string) bool {
	for _, str := range roles {
		if str == "admin" || str == "master" ||
			strings.Contains(str, ":admin") {
			return true
		}
	}
	return false
}

func Retry(maxTries int, f func() error) error {
	numTries := 0
	return doRetry(numTries, maxTries, f)
}

func doRetry(numTries int, maxTries int, f func() error) error {
	err := f()
	if err != nil && numTries < maxTries {
		numTries += 1
		return doRetry(numTries, maxTries, f)
	}
	return err
}

//Adds quotes to a string
func QuotifyString(str string) string {
	return "\"" + str + "\""
}

// Adds quotes to strings in an array,
// used for couchdb url parameters
func ApplyQuotes(strList []string) {
	for i, v := range strList {
		if v != "{}" {
			strList[i] = QuotifyString(v)
		}
	}
}

/* Parses an array in a query string.
/*
/* Returns an array
*/
func ParseArrayParam(str string) ([]string, error) {
	//The first and last characters should be '[' and ']'
	if str[0] != '[' || str[len(str)-1] != ']' {
		return []string{}, errors.New("Malformed array string.")
	}
	strArray := strings.Split(str[1:len(str)-1], ",")
	return strArray, nil
}

func GenHashString(data string) string {
	now := time.Now().UTC()
	mac := hmac.New(sha1.New, []byte(now.String()))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}

//Generates a random token
func GenToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

//Encodes a struct into json and returns an io.Reader
func EncodeJsonData(data interface{}) (io.Reader, int, error) {
	if data == nil {
		return nil, 0, nil
	}
	if buf, err := json.Marshal(&data); err != nil {
		return nil, 0, err
	} else {
		return bytes.NewReader(buf), len(buf), nil
	}
}

//Decodes a json response body and returns a struct
func DecodeJsonData(body io.Reader, result interface{}) error {
	decoder := json.NewDecoder(body)
	return decoder.Decode(result)
}

func EncodeBase64Url(data []byte) string {
	b64 := base64.StdEncoding.EncodeToString(data)
	b64 = strings.Replace(b64, "+", "-", -1)
	b64 = strings.Replace(b64, "/", "_", -1)
	b64 = strings.Replace(b64, "=", "", -1)
	return b64
}

//Returns a minimum of two integers
func MinInt(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

//Returns a maximum of two integers
func MaxInt(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

//Returns the Absolute value of an int
func AbsInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

// Krand，随机数（位数，类型)
/**
* size 随机码的位数
* kind 0    // 纯数字
       1    // 小写字母
       2    // 大写字母
       3    // 数字、大小写字母
*/
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

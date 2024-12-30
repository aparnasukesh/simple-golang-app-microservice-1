package user

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func ExtractErrorMessage(err error) string {
	if err == nil {
		return ""
	}
	errMsg := err.Error()
	if index := strings.Index(errMsg, "desc = "); index != -1 {
		return errMsg[index+len("desc = "):]
	}
	return errMsg
}

func HashPassword(password string) string {
	data := []byte(password)
	password = fmt.Sprintf("%x", md5.Sum(data))
	return password
}

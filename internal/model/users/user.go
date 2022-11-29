package users

import (
	"fmt"
)

type User struct {
	UserID   int64  `json:"userId,omitempty"`
	UserName string `json:"userName,omitempty"`
	FullName string `json:"fullName"`
}

func (u *User) Mention() string {
	if u == nil {
		return ""
	}
	if u.UserName != "" {
		return "@" + u.UserName
	}

	name := u.FullName

	if u.UserID != 0 {
		return fmt.Sprintf("[%s](tg://user?id=%o)", name, u.UserID)
	}

	return name
}

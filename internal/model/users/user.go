package users

type User struct {
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

	return u.FullName
}

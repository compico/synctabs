package booklink

import (
	"crypto/sha256"
	"encoding/json"
	"net/mail"
	"strings"
)

type User struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func NewUser(username string, password string, email string) (*User, error){
	user := new(User)
	user.Username = username
	user.Nickname = strings.ToUpper(username)
	switch emailValidation(email) {
		case false:
			return nil, EmailNotValid
		case true:
			user.Email = email
	}
	user.AddPassword(password)
	return user, nil
}

func (user *User) AddPassword(password string) {
	h := sha256.New()
	h.Write([]byte(password))
	user.Password = string(h.Sum(nil))
}

func (user *User) ToJson() ([]byte, error) {
	data, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func emailValidation(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

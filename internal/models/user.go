package models

import (
	"encoding/hex"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidPwdInput        = errors.New("possible empty password")
	ErrUsernameRequired       = errors.New("username required")
	ErrPwdRequired            = errors.New("password required")
	ErrAtLeastOneRoleRequired = errors.New("at least one role required")
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Roles    []Role `json:"role"`
}

func (u *User) HasRole(role Role) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}

	return false
}

func (u *User) Validate() error {
	if u.Username == "" {
		return ErrUsernameRequired
	}

	if u.Password == "" {
		return ErrPwdRequired
	}

	if len(u.Roles) == 0 {
		return ErrAtLeastOneRoleRequired
	}

	return nil
}

func (u *User) ParsePwd() error {
	const cost = 14
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), cost)
	u.Password = hex.EncodeToString(hash)

	return nil
}

func (u *User) CheckPwd(pwd string) bool {
	hash, _ := hex.DecodeString(u.Password)
	err := bcrypt.CompareHashAndPassword(hash, []byte(pwd))
	return err == nil
}

type Role string

const (
	Student Role = "STUDENT"
	Teacher Role = "TEACHER"
	Admin   Role = "ADMIN"
)

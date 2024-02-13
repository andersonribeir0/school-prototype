package models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
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
	hash := sha256.New()
	n, err := hash.Write([]byte(u.Password))
	if n == 0 || err != nil {
		return ErrInvalidPwdInput
	}

	hashBytes := hash.Sum(nil)
	u.Password = hex.EncodeToString(hashBytes)

	return nil
}

type Role string

const (
	Student Role = "STUDENT"
	Teacher Role = "TEACHER"
	Admin   Role = "ADMIN"
)

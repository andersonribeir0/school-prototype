package services

import (
	"context"
	"errors"

	"github.com/andersonribeir0/school-prototype/internal/db"
	"github.com/andersonribeir0/school-prototype/internal/logger"
	"github.com/andersonribeir0/school-prototype/internal/models"
)

var (
	ErrOnUserInsertion  = errors.New("user is already in the insert operation")
	ErrInvalidUserOrPwd = errors.New("invalid user or password")
	ErrUserInfo         = errors.New("it was not possible to get user information")
	ErrUserNotFound     = errors.New("user not found")
)

type UserSvc interface {
	AddUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) *models.User
	VerifyPwd(ctx context.Context, username string, pwd string) (*models.User, error)
}

type User struct {
	logger   logger.Logger
	userRepo db.UserRepository
}

func NewUserSvc(logger logger.Logger, userRepo db.UserRepository) *User {
	return &User{logger: logger, userRepo: userRepo}
}

func (u *User) AddUser(ctx context.Context, user *models.User) error {
	if err := user.ParsePwd(); err != nil {
		return err
	}

	id, err := u.userRepo.InsertUser(ctx, user)
	if err != nil {
		u.logger.Error("AddUser error inserting user " + err.Error())
		return ErrOnUserInsertion
	}

	u.logger.Info("used inserted " + id)
	return err
}

func (u *User) GetUserById(ctx context.Context, id string) *models.User {
	user, err := u.userRepo.GetUserById(ctx, id)
	if err != nil {
		u.logger.Error("GetUserById svc error: " + err.Error())
	}

	return user
}

func (u *User) VerifyPwd(ctx context.Context, username string, password string) (*models.User, error) {
	user := &models.User{
		Username: username,
		Password: password,
	}

	user, err := u.userRepo.GetUserByUsername(ctx, user.Username)
	if err != nil {
		return nil, ErrUserInfo
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	if !user.CheckPwd(password) {
		return nil, ErrInvalidUserOrPwd
	}

	return user, nil
}

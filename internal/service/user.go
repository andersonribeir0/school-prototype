package service

import (
	"context"

	"github.com/andersonribeir0/school-prototype/internal"
	"github.com/andersonribeir0/school-prototype/internal/db"
	"github.com/andersonribeir0/school-prototype/internal/logger"
)

type UserSvc interface {
	AddUser(ctx context.Context, user *internal.User) error
	GetUserById(ctx context.Context, id string) *internal.User
}

type User struct {
	logger   logger.Logger
	userRepo db.UserRepository
}

func NewUserSvc(logger logger.Logger, userRepo db.UserRepository) *User {
	return &User{logger: logger, userRepo: userRepo}
}

func (u *User) AddUser(ctx context.Context, user *internal.User) error {
	id, err := u.userRepo.InsertUser(ctx, user)
	u.logger.Info("used added " + id)
	return err
}

func (u *User) GetUserById(ctx context.Context, id string) *internal.User {
	user, err := u.userRepo.GetUserById(ctx, id)
	if err != nil {
		u.logger.Error("GetUserById svc error: " + err.Error())
	}

	return user
}

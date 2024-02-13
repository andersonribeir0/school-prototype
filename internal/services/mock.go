package services

import (
	"context"

	"github.com/andersonribeir0/school-prototype/internal/models"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) InsertUser(ctx context.Context, user *models.User) (string, error) {
	args := m.Called(ctx, user)
	return args.String(0), args.Error(1)
}

func (m *UserRepositoryMock) GetUserById(ctx context.Context, id string) (*models.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *UserRepositoryMock) GetUserByUsernameAndPassword(ctx context.Context, username, password string) (*models.User, error) {
	args := m.Called(ctx, username, password)
	return args.Get(0).(*models.User), args.Error(1)
}

type LoggerMock struct {
	mock.Mock
}

func (m *LoggerMock) Info(msg string, keysAndValues ...any) {
	m.Called(msg, keysAndValues)
}

func (m *LoggerMock) Warn(msg string, keysAndValues ...any) {
	m.Called(msg, keysAndValues)
}

func (m *LoggerMock) Debug(msg string, keysAndValues ...any) {
	m.Called(msg, keysAndValues)
}

func (m *LoggerMock) Error(msg string, keysAndValues ...any) {
	m.Called(msg, keysAndValues)
}

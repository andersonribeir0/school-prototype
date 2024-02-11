package service

import (
	"context"
	"testing"

	"github.com/andersonribeir0/school-prototype/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Info(msg string, args ...any) {
	m.Called(append([]any{msg}, args...)...)
}

func (m *MockLogger) Warn(msg string, args ...any) {
	m.Called(append([]any{msg}, args...)...)
}

func (m *MockLogger) Debug(msg string, args ...any) {
	m.Called(append([]any{msg}, args...)...)
}

func (m *MockLogger) Error(msg string, args ...any) {
	m.Called(append([]any{msg}, args...)...)
}

type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) InsertUser(ctx context.Context, user *internal.User) (string, error) {
	args := m.Called(ctx, user)
	return args.String(0), args.Error(1)
}

func (m *MockDatabase) GetUserById(ctx context.Context, id string) (*internal.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*internal.User), args.Error(1)
}

func TestAddUser(t *testing.T) {
	ctx := context.TODO()
	mockLogger := new(MockLogger)
	mockDb := new(MockDatabase)
	userSvc := NewUserSvc(mockLogger, mockDb)

	mockUser := &internal.User{ID: "testID", Name: "John Doe"}

	mockLogger.On("Info", mock.Anything).Return()
	mockDb.On("InsertUser", ctx, mockUser).Return(mockUser.ID, nil)

	err := userSvc.AddUser(ctx, mockUser)

	assert.NoError(t, err)
	mockLogger.AssertCalled(t, "Info", "used added "+mockUser.ID)
	mockDb.AssertExpectations(t)
}

func TestGetUserById(t *testing.T) {
	ctx := context.TODO()
	mockLogger := new(MockLogger)
	mockDb := new(MockDatabase)
	userSvc := NewUserSvc(mockLogger, mockDb)

	mockUser := &internal.User{ID: "testID", Name: "John Doe"}

	mockLogger.On("Error", mock.Anything).Return()
	mockDb.On("GetUserById", ctx, mockUser.ID).Return(mockUser, nil)

	result := userSvc.GetUserById(ctx, mockUser.ID)

	assert.Equal(t, mockUser, result)
	mockDb.AssertExpectations(t)
	mockLogger.AssertNotCalled(t, "Error", mock.Anything)
}

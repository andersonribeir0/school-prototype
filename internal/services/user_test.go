// userSvc_test.go

package services

import (
	"context"
	"testing"

	"github.com/andersonribeir0/school-prototype/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddUser(t *testing.T) {
	loggerMock := new(LoggerMock)
	userRepoMock := new(UserRepositoryMock)
	userSvc := NewUserSvc(loggerMock, userRepoMock)

	testUser := &models.User{
		Username: "testUser",
		Password: "testPass",
	}

	loggerMock.On("Info", mock.Anything, mock.Anything).Once()
	userRepoMock.On("InsertUser", mock.Anything, mock.AnythingOfType("*models.User")).Return(mock.Anything, nil)

	err := userSvc.AddUser(context.Background(), testUser)

	assert.NoError(t, err)
	userRepoMock.AssertExpectations(t)
}

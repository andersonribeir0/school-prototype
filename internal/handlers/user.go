package handlers

import (
	"context"
	"net/http"

	"github.com/andersonribeir0/school-prototype/internal"
	"github.com/labstack/echo/v4"
)

func (a *API) AddUserHandler(c echo.Context) error {
	err := a.userSvc.AddUser(context.Background(), &internal.User{
		Name: "Foo",
		Role: []internal.Role{internal.Student},
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "ok")
}

func (a *API) GetUserHandler(c echo.Context) error {
	user := a.userSvc.GetUserById(context.Background(), c.Param("id"))

	return c.JSON(http.StatusOK, user)
}

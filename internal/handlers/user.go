package handlers

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/andersonribeir0/school-prototype/internal/models"
	"github.com/andersonribeir0/school-prototype/internal/views/user"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (a *API) AddUserHandler(c echo.Context) error {
	var req AddUserRequest

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusOK, &Response{
			Body:    nil,
			Message: ErrInvalidBody.Error(),
			Success: false,
		})
	}

	err = a.userSvc.AddUser(c.Request().Context(), &models.User{
		Username: req.Username,
		Password: req.Password,
		Roles:    req.Roles,
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, &Response{
		Body:    nil,
		Message: "OK",
		Success: true,
	})
}

func (a *API) LoginHandler(c echo.Context) error {
	var loginReq LoginRequest
	if err := c.Bind(&loginReq); err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Body:    nil,
			Message: ErrInvalidBody.Error(),
			Success: false,
		})
	}

	userResp, err := a.userSvc.VerifyPwd(c.Request().Context(), loginReq.Username, loginReq.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, &Response{
			Body:    nil,
			Message: err.Error(),
			Success: false,
		})
	}

	claims := &jwtCustomClaims{
		loginReq.Username,
		userResp.Roles,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &Response{
		Body:    map[string]any{"token": t},
		Message: "",
		Success: true,
	})
}

func (a *API) GetUserHandler(c echo.Context) error {
	user := a.userSvc.GetUserById(context.Background(), c.Param("id"))

	return c.JSON(http.StatusOK, user)
}

func (a *API) GetMeHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	return c.JSON(http.StatusOK, echo.Map{
		"username": claims.Username,
		"roles":    claims.Roles,
	})
}

func (a *API) ShowMeHandler(c echo.Context) error {
	return render(c, user.Show())
}

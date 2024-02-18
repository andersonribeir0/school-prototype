package handlers

import (
	"context"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/andersonribeir0/school-prototype/internal/logger"
	"github.com/andersonribeir0/school-prototype/internal/models"
	"github.com/andersonribeir0/school-prototype/internal/services"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo    *echo.Echo
	logger  logger.Logger
	userSvc services.UserSvc
}

type jwtCustomClaims struct {
	Username string        `json:"username"`
	Roles    []models.Role `json:"roles"`
	jwt.RegisteredClaims
}

func NewAPI(logger logger.Logger, userSvc services.UserSvc) *API {
	e := echo.New()

	api := &API{Echo: e, logger: logger, userSvc: userSvc}

	api.Echo.POST("/login", api.LoginHandler)
	api.Echo.POST("/user", api.AddUserHandler)
	r := api.Echo.Group("/auth")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SIGNING_KEY")),
	}
	r.Use(echojwt.WithConfig(config))
	r.Use(withUser)
	r.GET("/", api.ShowMeHandler)
	r.GET("/me", api.GetMeHandler)

	return api
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		loggerUser, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "unauthorized",
			})
		}
		claims := loggerUser.Claims.(*jwtCustomClaims)
		ctx := context.WithValue(c.Request().Context(), "username", claims.Username)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

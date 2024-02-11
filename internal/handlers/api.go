package handlers

import (
	"github.com/andersonribeir0/school-prototype/internal/logger"
	"github.com/andersonribeir0/school-prototype/internal/service"
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo    *echo.Echo
	logger  logger.Logger
	userSvc service.UserSvc
}

func NewAPI(logger logger.Logger, userSvc service.UserSvc) *API {
	e := echo.New()

	return &API{Echo: e, logger: logger, userSvc: userSvc}
}

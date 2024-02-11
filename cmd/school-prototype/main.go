package main

import (
	"log/slog"

	"github.com/andersonribeir0/school-prototype/internal/db"
	"github.com/andersonribeir0/school-prototype/internal/handlers"
	"github.com/andersonribeir0/school-prototype/internal/service"
)

func main() {
	logger := slog.Default()
	logger.Info("Starting application...")

	database := db.NewDatabase(logger)
	userSvc := service.NewUserSvc(logger, database)
	api := handlers.NewAPI(logger, userSvc)

	api.Echo.POST("/user", api.AddUserHandler)
	api.Echo.GET("/user/:id", api.GetUserHandler)

	api.Echo.Start(":8080")
}

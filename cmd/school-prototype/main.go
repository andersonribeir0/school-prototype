package main

import (
	"log/slog"

	"github.com/andersonribeir0/school-prototype/internal/db"
	"github.com/andersonribeir0/school-prototype/internal/handlers"
	"github.com/andersonribeir0/school-prototype/internal/services"
)

func main() {
	logger := slog.Default()
	logger.Info("Starting application...")

	database := db.NewDatabase(logger)
	userSvc := services.NewUserSvc(logger, database)
	api := handlers.NewAPI(logger, userSvc)

	api.Echo.Start(":8080")
}

package handlers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

var datasource *sqlx.DB

func SetupHandlers(app *fiber.App, db *sqlx.DB) {
	datasource = db
	
	api := app.Group("/api/" + os.Getenv("API_VERSION"))
	
	pingHandler(api)
	usersHandler(api)
}

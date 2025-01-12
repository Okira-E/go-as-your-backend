package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/org/example/internal/server_health"
	"github.com/org/example/internal/system_users"
)

func main() {
	// Load the environment variables.
	err := setupEnv()
	if err != nil {
		log.Fatalln("Error loading .env file. " + err.Error())
	}

	// Setup the daatbase connection using sqlx.
	db, err := setupDatabase()
	if err != nil {
		log.Fatalf("Error connecting to the database. %s\n", err.Error())
	}
	defer db.Close()

	app := fiber.New()

	// Setup CORS.
	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept",
		},
	))

	// Recover from panics anywhere.
	app.Use(recover.New())

	// If you want to setup a project-wide basice auth on every endpoint.
	// You could also do bearer auth here.
	/*
		app.Use(basicauth.New(basicauth.Config{
		    SystemUsers: map[string]string{
		        "john":  "doe",
		        "admin": "123456",
		    },
		}))
	*/

	// Setup logging
	mode := os.Getenv("ENV")
	if mode == "debug" {
		app.Use(logger.New(logger.Config{
			Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
		}))
	} else if mode == "prod" {
		// Setup optional writing to persitstant logs.
	}
	fmt.Printf("Running in %s mode.\n", mode)

	setupRoutes(app, db)

	// Run server.
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("Environment variable PORT is not set.")
	}

	host := os.Getenv("HOST")
	if host == "" {
		log.Fatalln("Environment variable HOST is not set.")
	}

	log.Fatalln(app.Listen(host + ":" + port))
}

func setupEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func setupDatabase() (*sqlx.DB, error) {
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbname, password))
	if err != nil {
		return nil, err
	}

	// Set the maximum number of open connections
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)   // Control idle connections
	db.SetConnMaxLifetime(0) // No limit, or set a duration like time.Minute * 5

	return db, nil
}

func setupRoutes(app *fiber.App, db *sqlx.DB) {
	version := os.Getenv("API_VERSION")
	if version == "" {
		log.Fatalln("API_VERSION environment variable is not set.")
	}

	api := app.Group("/api/" + version)

	server_health.SetupHandlers(api)
	
	system_users.InitConfig(db)
	system_users.SetupHandlers(api)
}

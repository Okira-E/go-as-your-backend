package handlers

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/org/example/internal/server/features"
	"github.com/org/example/internal/server/models"
	"github.com/org/example/internal/server/utils"
)

func usersHandler(api fiber.Router) {
	path := "/users"

	api.Get(path, getUsers)
	api.Post(path, registerUser)
	api.Delete(path+"/:id", deleteUser)
}

// Example usage: "http://localhost:3200?filter={\"where\": {\"id\": \"123-456-789\"}, \"order_by\": [\"created_by\"]}&limit=50&offset=10"
// Without the filter param, it returns all the users.
func getUsers(c *fiber.Ctx) error {
	path, err := url.Parse(c.OriginalURL())
	if err != nil {
		log.Fatalf("Failed to parse the URL. %s", err)
	}

	queryParams := path.Query()

	usersDto, err := features.GetAllUsers(datasource, queryParams)
	if err != nil {
		return utils.Err(c, 500, "Error fetching users. "+err.Error(), nil)
	}

	return utils.Ok(c, 200, "", usersDto)
}

func registerUser(c *fiber.Ctx) error {
	payload := struct {
		models.UsersDto

		Password string `json:"password" validate:"required"`
	}{}

	err := json.Unmarshal(c.Body(), &payload)
	if err != nil {
		return utils.Err(c, 400, "Error parsing user. "+err.Error(), nil)
	}

	// Validate the request body using the struct Validator tags.
	v := validator.New(validator.WithRequiredStructEnabled())
	err = v.Struct(payload)
	if err != nil {
		return utils.Err(c, 400, "Request body is not valid. "+err.Error(), nil)
	}

	var userDto models.UsersDto
	copier.Copy(&userDto, &payload.UsersDto)

	err = features.CreateUser(datasource, payload.UsersDto, payload.Password)
	if err != nil {
		return utils.Err(c, 500, "Error creating user. "+err.Error(), nil)
	}

	return utils.Ok(c, 201, "User created successfully.", nil)
}

func deleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.Err(c, 400, "User ID is required.", nil)
	}

	err := features.DeleteUser(datasource, id)
	if err != nil {
		return utils.Err(c, 500, "Error deleting user. "+err.Error(), nil)
	}

	return utils.Ok(c, 200, "User deleted successfully.", nil)
}

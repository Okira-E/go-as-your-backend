package system_users

import (
	"encoding/json"
	"github.com/org/example/internal/utils"
	"log"
	"net/url"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

func SetupHandlers(api fiber.Router) {
	api = api.Group("/users")

	api.Get("/", getUsers)
	api.Post("/", registerUser)
	api.Delete("/:id", deleteUser)
}

// Example usage: "http://localhost:3200?filter={\"where\": {\"id\": \"123-456-789\"}, \"order_by\": [\"created_by\"]}&limit=50&offset=10"
// Without the filter param, it returns all the users.
func getUsers(c *fiber.Ctx) error {
	path, err := url.Parse(c.OriginalURL())
	if err != nil {
		log.Fatalf("Failed to parse the URL. %s", err)
	}

	queryParams := path.Query()

	usersDto, err := GetAllUsers(config.datasource, queryParams)
	if err != nil {
		return utils.Err(c, 500, "Error fetching users. "+err.Error(), nil)
	}

	return utils.Ok(c, 200, "", usersDto)
}

func registerUser(c *fiber.Ctx) error {
	payload := struct {
		SystemUsersDto

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

	var userDto SystemUsersDto
	err = copier.Copy(&userDto, &payload.SystemUsersDto)
	if err != nil {
		return utils.Err(c, 500, "Encountered an error while constructing the object to save. "+err.Error(), nil)
	}

	insertedUser, status, err := CreateUser(config.datasource, payload.SystemUsersDto, payload.Password)
	if err != nil {
		return utils.Err(c, status, "Error creating user. "+err.Error(), nil)
	}

	return utils.Ok(c, status, "User created successfully.", insertedUser)
}

func deleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.Err(c, 400, "Entity ID is required.", nil)
	}

	err := DeleteUser(config.datasource, id)
	if err != nil {
		return utils.Err(c, 500, "Error deleting entity. "+err.Error(), nil)
	}

	return utils.Ok(c, 200, "Entity deleted successfully.", nil)
}

package features

import (
	"fmt"
	"net/url"

	"github.com/jinzhu/copier"
	"github.com/jmoiron/sqlx"
	"github.com/org/example/internal/server/models"
	"github.com/org/example/internal/server/utils"
)

func GetAllUsers(datasource *sqlx.DB, queryParams url.Values) ([]models.UsersDto, error) {
	limit, offset, err := utils.ParseLimitAndOffset(queryParams)
	if err != nil {
		return nil, fmt.Errorf("Error parsing the limit & offset."+err.Error())
	}

	filter, err := utils.ParseFilterClause(queryParams)
	if err != nil {
		return nil, fmt.Errorf("Error parsing the fiter query."+err.Error())
	}

	query, args := utils.ParseQuery("users", limit, offset, filter)

	rows, err := datasource.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	
	users := []models.Users{}
	for rows.Next() {
		var user models.Users
		err = rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	var usersDto []models.UsersDto
	copier.Copy(&usersDto, &users)
	
	return usersDto, nil
}

func CreateUser(datasource *sqlx.DB, userDto models.UsersDto, password string) error {
	// Pretend we hashed the password given.
	// ...
	
	_, err := datasource.Exec("INSERT INTO users (email, username, first_name, last_name, password_hash) VALUES ($1, $2, $3, $4, $5)", userDto.Email, userDto.Username, userDto.FirstName, userDto.LastName, password)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(datasource *sqlx.DB, id string) error {
	_, err := datasource.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
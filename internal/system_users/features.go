package system_users

import (
	"fmt"
	"net/url"

	"github.com/jmoiron/sqlx"
	"github.com/org/example/internal/utils"
)

func GetAllUsers(datasource *sqlx.DB, queryParams url.Values) ([]SystemUsersDto, error) {
	limit, offset, err := utils.ParseLimitAndOffset(queryParams)
	if err != nil {
		return nil, fmt.Errorf("Error parsing the limit & offset." + err.Error())
	}

	filter, err := utils.ParseFilterClause(queryParams)
	if err != nil {
		return nil, fmt.Errorf("Error parsing the fiter query." + err.Error())
	}

	query, args := utils.MakeSelectStmt("system_users", limit, offset, filter)

	rows, err := datasource.Queryx(query, args...)
	if err != nil {
		return nil, err
	}

	users := []SystemUsers{}
	for rows.Next() {
		var user SystemUsers
		err = rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	var usersDto []SystemUsersDto
	for _, user := range users {
		usersDto = append(usersDto, user.ToSystemUsersDto())
	}

	return usersDto, nil
}

func CreateUser(datasource *sqlx.DB, userDto SystemUsersDto, password string) (SystemUsersDto, int, error) {
	// Pretend we hashed the password given.
	hashedPassword := password

	user := userDto.ToSystemUsers()

	// Default values.
	user.IsActive = true
	user.PasswordHash = hashedPassword

	stmt, values, err := utils.MakeInsertStmt(user)
	if err != nil {
		return SystemUsersDto{}, 500, fmt.Errorf("Error generating Insert statement for entity. %s", err)
	}

	stmt = fmt.Sprintf("%s RETURNING *", stmt)
	var insertedUser SystemUsers
	err = datasource.QueryRowx(stmt, values...).StructScan(&insertedUser)
	if err != nil {
		return SystemUsersDto{}, 400, fmt.Errorf("Error while inserting values in the database. %s", err)
	}

	insertedUserDto := insertedUser.ToSystemUsersDto()

	return insertedUserDto, 201, nil
}

func DeleteUser(datasource *sqlx.DB, id string) error {
	_, err := datasource.Exec("DELETE FROM system_users WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

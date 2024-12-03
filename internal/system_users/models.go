package system_users

type SystemUsers struct {
	ID           int    `db:"id"              json:"id"`
	Username     string `db:"username"        json:"username"`
	Email        string `db:"email"           json:"email"`
	PasswordHash string `db:"password_hash"   json:"password_hash"`
	FirstName    string `db:"first_name"      json:"first_name"`
	LastName     string `db:"last_name"       json:"last_name"`
	IsActive     bool   `db:"is_active"       json:"is_active"`
	CreatedAt    string `db:"created_at"      json:"created_at"`
	UpdatedAt    string `db:"updated_at"      json:"updated_at"`
}

func (systemUser *SystemUsers) ToSystemUsersDto() SystemUsersDto {
	return SystemUsersDto{
		ID:        systemUser.ID,
		Username:  systemUser.Username,
		Email:     systemUser.Email,
		FirstName: systemUser.FirstName,
		LastName:  systemUser.LastName,
		IsActive:  systemUser.IsActive,
		CreatedAt: systemUser.CreatedAt,
		UpdatedAt: systemUser.UpdatedAt,
	}
}

type SystemUsersDto struct {
	ID        int    `db:"id"              json:"id"`
	Username  string `db:"username"        json:"username"               validate:"required"`
	Email     string `db:"email"           json:"email"                  validate:"required"`
	FirstName string `db:"first_name"      json:"first_name"             validate:"required"`
	LastName  string `db:"last_name"       json:"last_name"              validate:"required"`
	IsActive  bool   `db:"is_active"       json:"is_active"`
	CreatedAt string `db:"created_at"      json:"created_at"`
	UpdatedAt string `db:"updated_at"      json:"updated_at"`
}

func (systemUserDto *SystemUsersDto) ToSystemUsers() SystemUsers {
	return SystemUsers{
		ID:           systemUserDto.ID,
		Username:     systemUserDto.Username,
		Email:        systemUserDto.Email,
		PasswordHash: "",
		FirstName:    systemUserDto.FirstName,
		LastName:     systemUserDto.LastName,
		IsActive:     systemUserDto.IsActive,
		CreatedAt:    systemUserDto.CreatedAt,
		UpdatedAt:    systemUserDto.UpdatedAt,
	}
}

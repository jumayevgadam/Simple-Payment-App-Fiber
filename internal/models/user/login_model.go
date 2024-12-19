package user

// LoginRequest model.
type LoginRequest struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}

type LoginResponse struct {
	UserID   int    `json:"userID"`
	RoleID   int    `json:"roleID"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	RoleType string `json:"roleType"`
	Username string `json:"username"`
}

type LoginResponseWithToken struct {
	LoginResponse
	Token string `json:"token"`
}

type LoginResponseData struct {
	UserID   int    `db:"id"`
	RoleID   int    `db:"role_id"`
	Name     string `db:"name"`
	Surname  string `db:"surname"`
	RoleType string `db:"role_type"`
	Username string `db:"username"`
	Password string `db:"password"`
}

func (l *LoginResponseData) ToServer() LoginResponse {
	return LoginResponse{
		UserID:   l.UserID,
		RoleID:   l.RoleID,
		Name:     l.Name,
		Surname:  l.Surname,
		RoleType: l.RoleType,
		Username: l.Username,
	}
}

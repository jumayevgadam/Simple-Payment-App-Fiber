package user

// UserWithTokens model is user details with token.
type UserWithTokens struct {
	Token string `json:"Token"`
}

type Details struct {
	ID       int    `db:"id"`
	RoleID   int    `db:"role_id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

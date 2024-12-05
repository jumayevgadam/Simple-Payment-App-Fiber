package user

type Details struct {
	ID       int    `db:"id"`
	RoleID   int    `db:"role_id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

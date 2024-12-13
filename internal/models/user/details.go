package user

type Details struct {
	ID       int    `db:"id"`
	RoleID   int    `db:"role_id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type StudentDTO struct {
	ID       int    `json:"id"`
	RoleID   int    `json:"roleID"`
	GroupID  int    `json:"groupID"`
	FullName string `json:"fullName"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type StudentDAO struct {
	ID       int    `db:"id"`
	RoleID   int    `db:"role_id"`
	GroupID  int    `db:"group_id"`
	FullName string `db:"full_name"`
	Username string `db:"username"`
	Password string `db:"password"`
}

func (s *StudentDAO) ToServer() *StudentDTO {
	return &StudentDTO{
		ID:       s.ID,
		RoleID:   s.RoleID,
		GroupID:  s.GroupID,
		FullName: s.FullName,
		Username: s.Username,
		Password: s.Password,
	}
}

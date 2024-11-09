package user

import "time"

// LoginReq model is
type SignUpReq struct {
	RoleID   int    `form:"role-id" json:"roleID" validate:"required"`
	GroupID  int    `form:"group-id" json:"groupID" validate:"required"`
	Name     string `form:"name" json:"name" validate:"required"`
	Surname  string `form:"surname" json:"surname" validate:"required"`
	UserName string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required,min=6"`
}

// LoginRes model is
type SignUpRes struct {
	RoleID   int    `db:"role_id"`
	GroupID  int    `db:"group_id"`
	Name     string `db:"name"`
	Surname  string `db:"surname"`
	UserName string `db:"username"`
	Password string `db:"password"`
}

// ToServer is
func (s *SignUpRes) ToServer() *SignUpReq {
	return &SignUpReq{
		RoleID:   s.RoleID,
		GroupID:  s.GroupID,
		Name:     s.Name,
		Surname:  s.Surname,
		UserName: s.UserName,
		Password: s.Password,
	}
}

// ToStorage is
func (s *SignUpReq) ToStorage() *SignUpRes {
	return &SignUpRes{
		RoleID:   s.RoleID,
		GroupID:  s.GroupID,
		Name:     s.Name,
		Surname:  s.Surname,
		UserName: s.UserName,
		Password: s.Password,
	}
}

// LoginReq model is
type LoginReq struct {
	Username string `form:"username" json:"userName" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

// LoginRes model is
type LoginRes struct {
	Username string `db:"username"`
	Password string `db:"password"`
}

// ToServer is
func (l *LoginRes) ToServer() *LoginReq {
	return &LoginReq{
		Username: l.Username,
		Password: l.Password,
	}
}

// ToStorage is
func (l *LoginReq) ToStorage() *LoginRes {
	return &LoginRes{
		Username: l.Username,
		Password: l.Password,
	}
}

// AllUserDAO model is
type AllUserDAO struct {
	ID        int       `db:"id"`
	RoleID    int       `db:"role_id"`
	GroupID   int       `db:"group_id"`
	Name      string    `db:"name"`
	Surname   string    `db:"surname"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// AllUserDTO model is
type AllUserDTO struct {
	ID       int    `json:"userID"`
	RoleID   int    `form:"role-id" json:"roleID"`
	GroupID  int    `form:"group-id" json:"groupID"`
	Name     string `form:"name" json:"name"`
	Surname  string `form:"surname" json:"surname"`
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

// ToStorage is
func (a *AllUserDTO) ToStorage() *AllUserDAO {
	return &AllUserDAO{
		ID:       a.ID,
		RoleID:   a.RoleID,
		GroupID:  a.GroupID,
		Name:     a.Name,
		Surname:  a.Surname,
		Username: a.Username,
		Password: a.Password,
	}
}

// ToServer is
func (a *AllUserDAO) ToServer() *AllUserDTO {
	return &AllUserDTO{
		ID:       a.ID,
		RoleID:   a.RoleID,
		GroupID:  a.GroupID,
		Name:     a.Name,
		Surname:  a.Surname,
		Username: a.Username,
		Password: a.Password,
	}
}

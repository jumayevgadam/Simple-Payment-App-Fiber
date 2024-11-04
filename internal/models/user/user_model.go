package user

// LoginReq model is
type SignUpReq struct {
	RoleID   int    `json:"roleID"`
	Name     string `form:"name" json:"name" validate:"required"`
	Surname  string `form:"surname" json:"surname" validate:"required"`
	UserName string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required,min=6"`
}

// LoginRes model is
type SignUpRes struct {
	RoleID   int    `db:"role_id"`
	Name     string `db:"name"`
	Surname  string `db:"surname"`
	UserName string `db:"username"`
	Password string `db:"password"`
}

// ToServer is
func (s *SignUpRes) ToServer() *SignUpReq {
	return &SignUpReq{
		RoleID:   s.RoleID,
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

// UserReq model is
type UserReq struct {
	RoleID   int    `form:"role-id" json:"roleID" validate:"required"`
	ClassID  int    `form:"class-id" json:"classID" validate:"required"`
	Name     string `form:"name" json:"name" validate:"required"`
	UserName string `form:"username" json:"userName" validate:"required,min=5"`
	Surname  string `form:"surname" json:"surName" validate:"required"`
	Password string `form:"password" json:"password" validate:"required,min=6"`
}

// UserRes model is
type UserRes struct {
	RoleID   int    `db:"role_id"`
	ClassID  int    `db:"class_id"`
	Name     string `db:"name"`
	UserName string `db:"username"`
	Surname  string `db:"surname"`
	Password string `db:"password"`
}

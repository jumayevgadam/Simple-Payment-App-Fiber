package user

import "time"

// LoginReq model is request model for processing request in handler layer.
type SignUpReq struct {
	Name     string `form:"name" json:"name" validate:"required"`
	Surname  string `form:"surname" json:"surname" validate:"required"`
	UserName string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required,min=6"`
	GroupID  *int   `form:"group-id,omitempty"`
}

// LoginRes model is db model.
type SignUpRes struct {
	RoleID   *int   `db:"role_id"`
	GroupID  *int   `db:"group_id"`
	Name     string `db:"name"`
	Surname  string `db:"surname"`
	UserName string `db:"username"`
	Password string `db:"password"`
}

// ToStorage method sends dto model to db storage.
func (s *SignUpReq) ToStorage() SignUpRes {
	return SignUpRes{
		GroupID:  s.GroupID,
		Name:     s.Name,
		Surname:  s.Surname,
		UserName: s.UserName,
		Password: s.Password,
	}
}

// LoginReq model is request model for processing request in handler layer.
type LoginReq struct {
	Username string `form:"username" json:"userName" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

// LoginRes model is response model which taken from DB.
type LoginRes struct {
	Username string `db:"username"`
	Password string `db:"password"`
}

// ToServer method sends response to the server.
func (l *LoginRes) ToServer() LoginReq {
	return LoginReq{
		Username: l.Username,
		Password: l.Password,
	}
}

// ToStorage method receives LoginReq model into db.
func (l *LoginReq) ToStorage() LoginRes {
	return LoginRes{
		Username: l.Username,
		Password: l.Password,
	}
}

// AllUserDAO model is data access object.
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

// AllUserDTO model is dto model.
type AllUserDTO struct {
	ID       int    `json:"userID"`
	RoleID   int    `form:"role-id" json:"roleID"`
	GroupID  int    `form:"group-id" json:"groupID"`
	Name     string `form:"name" json:"name"`
	Surname  string `form:"surname" json:"surname"`
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

// ToStorage method for receiving AllUserDTO into database.
func (a *AllUserDTO) ToStorage() AllUserDAO {
	return AllUserDAO{
		ID:       a.ID,
		RoleID:   a.RoleID,
		GroupID:  a.GroupID,
		Name:     a.Name,
		Surname:  a.Surname,
		Username: a.Username,
		Password: a.Password,
	}
}

// ToServer method sends AllUserDAO model to server.
func (a *AllUserDAO) ToServer() AllUserDTO {
	return AllUserDTO{
		ID:       a.ID,
		RoleID:   a.RoleID,
		GroupID:  a.GroupID,
		Name:     a.Name,
		Surname:  a.Surname,
		Username: a.Username,
		Password: a.Password,
	}
}

// StudentInfo model for Payment Model.
type StudentInfoData struct {
	CourseYear int    `db:"course_year"`
	FullName   string `db:"full_name"`
	Username   string `db:"username"`
	GroupCode  string `db:"group_code"`
}

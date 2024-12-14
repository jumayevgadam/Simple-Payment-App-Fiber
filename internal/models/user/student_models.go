package user

import "time"

// Request model.
type Request struct {
	GroupID  int    `form:"groupID" validate:"required"`
	Name     string `form:"name" validate:"required"`
	Surname  string `form:"surname" validate:"required"`
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}

// Response model.
type Response struct {
	RoleID   int    `db:"role_id"`
	GroupID  int    `db:"group_id"`
	Name     string `db:"name"`
	Surname  string `db:"surname"`
	Username string `db:"username"`
	Password string `db:"password"`
}

// ToPsqlDBStorage.
func (r *Request) ToPsqlDBStorage() Response {
	return Response{
		GroupID:  r.GroupID,
		Name:     r.Name,
		Surname:  r.Surname,
		Username: r.Username,
		Password: r.Password,
	}
}

// Student model.
type Student struct {
	ID        int       `json:"id"`
	GroupID   int       `json:"groupID"`
	RoleID    int       `json:"roleID"`
	Name      string    `json:"studentName"`
	Surname   string    `json:"studentSurname"`
	Username  string    `json:"studentUsername"`
	Password  string    `json:"studentPassword"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// StudentData model.
type StudentData struct {
	ID        int       `db:"id"`
	GroupID   int       `db:"group_id"`
	RoleID    int       `db:"role_id"`
	Name      string    `db:"name"`
	Surname   string    `db:"surname"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// ToServer func.
func (s *StudentData) ToServer() *Student {
	return &Student{
		ID:        s.ID,
		GroupID:   s.GroupID,
		RoleID:    s.RoleID,
		Name:      s.Name,
		Surname:   s.Surname,
		Username:  s.Username,
		Password:  s.Password,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

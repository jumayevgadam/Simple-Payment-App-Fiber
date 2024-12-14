package user

import (
	"time"

	"github.com/jumayevgadam/tsu-toleg/pkg/constants"
)

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

type StudentDataByGroupID struct {
	ID        int       `db:"id"`
	RoleID    int       `db:"role_id"`
	GroupID   int       `db:"group_id"`
	FullName  string    `db:"full_name"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type StudentResGroupID struct {
	ID        int       `json:"id"`
	RoleID    int       `json:"roleID"`
	GroupID   int       `json:"groupID"`
	FullName  string    `json:"fullName"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (s *StudentDataByGroupID) ToServer() *StudentResGroupID {
	return &StudentResGroupID{
		ID:        s.ID,
		RoleID:    s.RoleID,
		GroupID:   s.GroupID,
		FullName:  s.FullName,
		Username:  s.Username,
		Password:  s.Password,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

type StudentUpdateRequest struct {
	GroupID  int    `form:"groupID"`
	Name     string `form:"name"`
	Surname  string `form:"surname"`
	Username string `form:"username"`
	Password string `form:"password"`
}

type StudentUpdateData struct {
	StudentID int    `db:"id"`
	GroupID   int    `db:"group_id"`
	Name      string `db:"name"`
	Surname   string `db:"surname"`
	Username  string `db:"username"`
	Password  string `db:"password"`
}

func (s *StudentUpdateRequest) ToPsqlDBStorage(studentID int) StudentUpdateData {
	return StudentUpdateData{
		StudentID: studentID,
		GroupID:   s.GroupID,
		Name:      s.Name,
		Surname:   s.Surname,
		Username:  s.Username,
		Password:  s.Password,
	}
}

func (s *StudentUpdateRequest) Validate() (string, error) {
	if s.GroupID == 0 && s.Name == "" && s.Surname == "" && s.Username == "" && s.Password == "" {
		return constants.NoUpdateResponse, nil
	}

	return "", nil
}

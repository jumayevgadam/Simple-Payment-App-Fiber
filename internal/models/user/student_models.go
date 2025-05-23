package user

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/pkg/constants"
)

// Request model.
type Request struct {
	GroupID  int    `form:"group-id" validate:"required"`
	Name     string `form:"name" validate:"required"`
	Surname  string `form:"surname" validate:"required"`
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required,gte=6"`
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
	Name      string    `db:"name"`
	Surname   string    `db:"surname"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type StudentResGroupID struct {
	ID        int       `json:"id"`
	RoleID    int       `json:"roleID"`
	GroupID   int       `json:"groupID"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
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
		Name:      s.Name,
		Surname:   s.Surname,
		Username:  s.Username,
		Password:  s.Password,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

type StudentUpdateRequest struct {
	GroupID  int    `form:"group-id"`
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

type AllStudentData struct {
	StudentID       int    `db:"student_id"`
	StudentName     string `db:"student_name"`
	StudentSurname  string `db:"student_surname"`
	StudentUsername string `db:"student_username"`
	RoleName        string `db:"role_name"`
	FacultyName     string `db:"faculty_name"`
	GroupCode       string `db:"group_code"`
	CourseYear      int    `db:"course_year"`
}

type AllStudentDTO struct {
	StudentID       int    `json:"studentID"`
	StudentName     string `json:"studentName"`
	StudentSurname  string `json:"studentSurname"`
	StudentUsername string `json:"studentUsername"`
	RoleName        string `json:"role"`
	FacultyName     string `json:"facultyName"`
	GroupCode       string `json:"groupCode"`
	CourseYear      int    `json:"courseYear"`
}

func (a *AllStudentData) ToServer() *AllStudentDTO {
	return &AllStudentDTO{
		StudentID:       a.StudentID,
		StudentName:     a.StudentName,
		StudentSurname:  a.StudentSurname,
		StudentUsername: a.StudentUsername,
		RoleName:        a.RoleName,
		FacultyName:     a.FacultyName,
		GroupCode:       a.GroupCode,
		CourseYear:      a.CourseYear,
	}
}

type FilterStudent struct {
	FacultyName    string
	GroupCode      string
	StudentName    string
	StudentSurname string
	PaymentStatus  string
}

func GetQueryParamsForFilterStudents(c *fiber.Ctx) FilterStudent {
	return FilterStudent{
		FacultyName:    c.Query("faculty-name", ""),
		GroupCode:      c.Query("group-code", ""),
		StudentName:    c.Query("student-name", ""),
		StudentSurname: c.Query("student-surname", ""),
		PaymentStatus:  c.Query("payment-status", ""),
	}
}

type StudentNameAndSurnameData struct {
	Name    string `db:"name"`
	Surname string `db:"surname"`
}

type StudentNameAndSurname struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (s *StudentNameAndSurnameData) ToServer() StudentNameAndSurname {
	return StudentNameAndSurname{
		Name:    s.Name,
		Surname: s.Surname,
	}
}

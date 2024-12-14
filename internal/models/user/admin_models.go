package user

import (
	"time"

	"github.com/jumayevgadam/tsu-toleg/pkg/constants"
)

// AdminRequest model.
type AdminRequest struct {
	Name     string `form:"name"`
	Surname  string `form:"surname"`
	Username string `form:"username"`
	Password string `form:"password"`
}

// AdminResponse model.
type AdminResponse struct {
	Name     string `db:"name"`
	Surname  string `db:"surname"`
	Username string `db:"username"`
	Password string `db:"password"`
}

// ToPsqlDBStorage.
func (a *AdminRequest) ToPsqlDBStorage() AdminResponse {
	return AdminResponse{
		Name:     a.Name,
		Surname:  a.Surname,
		Username: a.Username,
		Password: a.Password,
	}
}

// AllAdminData model.
type AdminData struct {
	ID        int       `db:"id"`
	RoleID    int       `db:"role_id"`
	Name      string    `db:"name"`
	Surname   string    `db:"surname"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Admin model.
type Admin struct {
	ID        int       `json:"id"`
	RoleID    int       `json:"role_id"`
	Name      string    `json:"adminName"`
	Surname   string    `json:"adminSurname"`
	Username  string    `json:"adminUsername"`
	Password  string    `json:"adminPassword"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ToServer func.
func (a *AdminData) ToServer() *Admin {
	return &Admin{
		ID:        a.ID,
		RoleID:    a.RoleID,
		Name:      a.Name,
		Surname:   a.Surname,
		Username:  a.Username,
		Password:  a.Password,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

// AdminUpdateRequest model.
type AdminUpdateRequest struct {
	Name     string `form:"name"`
	Surname  string `form:"surname"`
	Username string `form:"username"`
	Password string `form:"password"`
}

type AdminUpdateData struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Surname  string `db:"surname"`
	UserName string `db:"username"`
	Password string `db:"password"`
}

func (a *AdminUpdateRequest) ToPsqlDBStorage(adminID int) AdminUpdateData {
	return AdminUpdateData{
		ID:       adminID,
		Name:     a.Name,
		Surname:  a.Surname,
		UserName: a.Username,
		Password: a.Password,
	}
}

func (a *AdminUpdateRequest) Validate() (string, error) {
	if a.Name == "" && a.Username == "" && a.Surname == "" && a.Password == "" {
		return constants.NoUpdateResponse, nil
	}

	return "", nil
}

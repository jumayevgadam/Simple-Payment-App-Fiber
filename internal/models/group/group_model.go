package group

import "time"

// DTO model is
type DTO struct {
	ID           int    `json:"classID"`
	DepartmentID int    `form:"department-id" json:"departmentID" validate:"required"`
	ClassCode    string `form:"class-code" json:"classCode" validate:"required"`
}

// DAO model is
type DAO struct {
	ID           int       `db:"id"`
	DepartmentID int       `db:"department_id"`
	ClassCode    string    `db:"code_name"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

// ToServer is
func (d *DAO) ToServer() *DTO {
	return &DTO{
		ID:           d.ID,
		DepartmentID: d.DepartmentID,
		ClassCode:    d.ClassCode,
	}
}

// ToStorage is
func (d *DTO) ToStorage() *DAO {
	return &DAO{
		ID:           d.ID,
		DepartmentID: d.DepartmentID,
		ClassCode:    d.ClassCode,
	}
}

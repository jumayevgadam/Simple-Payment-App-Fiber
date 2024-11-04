package faculty

import "time"

// We use in this model type
// DTO and DAO models

// DTO is
type DTO struct {
	ID   int    `json:"facultyID"`
	Name string `form:"faculty-name" json:"facultyName" validate:"required"`
}

// DAO is
type DAO struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// ToServer is
func (d *DAO) ToServer() *DTO {
	return &DTO{
		ID:   d.ID,
		Name: d.Name,
	}
}

// ToStorage is
func (d *DTO) ToStorage() *DAO {
	return &DAO{
		ID:   d.ID,
		Name: d.Name,
	}
}

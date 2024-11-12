package role

// DTO model is data transfer object: =>
type DTO struct {
	ID       int    `json:"roleID"`
	RoleName string `form:"role-name" json:"roleName" validate:"required"`
}

// DAO model is data access object: =>
type DAO struct {
	ID       int    `db:"id"`
	RoleName string `db:"name"`
}

// ToServer is
func (d *DAO) ToServer() DTO {
	return DTO{
		ID:       d.ID,
		RoleName: d.RoleName,
	}
}

// ToStorage is
func (d *DTO) ToStorage() DAO {
	return DAO{
		ID:       d.ID,
		RoleName: d.RoleName,
	}
}

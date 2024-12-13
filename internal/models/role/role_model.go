package role

// DTO model is data transfer object: =>.
type DTO struct {
	ID       int    `json:"roleID"`
	RoleName string `form:"role" json:"role"`
}

// DAO model is data access object: =>.
type DAO struct {
	ID       int    `db:"id"`
	RoleName string `db:"role"`
}

// ToServer method sends DAO model to server.
func (d *DAO) ToServer() DTO {
	return DTO{
		ID:       d.ID,
		RoleName: d.RoleName,
	}
}

// ToStorage method sends DTO to DAO(saves it into database).
func (d *DTO) ToStorage() DAO {
	return DAO{
		ID:       d.ID,
		RoleName: d.RoleName,
	}
}

package faculty

// We use in this model type.
// DTO and DAO models.

// DTO model is data transfer object.
type DTO struct {
	ID   int    `json:"facultyID"`
	Name string `form:"faculty-name" json:"facultyName"`
	Code string `form:"faculty-code" json:"faculty-code"`
}

// DAO model is data access object.
type DAO struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Code string `db:"code"`
}

// ToServer method sends DAO model to Server.
func (d *DAO) ToServer() DTO {
	return DTO{
		ID:   d.ID,
		Name: d.Name,
		Code: d.Code,
	}
}

// ToStorage model receives DTO model and save it to the database.
func (d *DTO) ToStorage() DAO {
	return DAO{
		ID:   d.ID,
		Name: d.Name,
		Code: d.Code,
	}
}

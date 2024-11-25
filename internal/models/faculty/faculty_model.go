package faculty

// We use in this model type.
// DTO and DAO models.

// DTO model is data transfer object.
type DTO struct {
	Name string `form:"faculty-name" json:"facultyName" validate:"required"`
	Code string `form:"faculty-code" json:"faculty-code" validate:"required"`
}

// DAO model is data access object.
type DAO struct {
	Name string `db:"name"`
	Code string `db:"code"`
}

// ToServer method sends DAO model to Server.
func (d *DAO) ToServer() *DTO {
	return &DTO{
		Name: d.Name,
		Code: d.Code,
	}
}

// ToStorage model receives DTO model and save it to the database.
func (d *DTO) ToStorage() *DAO {
	return &DAO{
		Name: d.Name,
		Code: d.Code,
	}
}

// Faculty model is DTO model.
type Faculty struct {
	ID   int    `json:"facultyID"`
	Name string `form:"faculty-name" json:"facultyName" validate:"required"`
	Code string `fomr:"faculty-code" json:"faculty-code" validate:"required"`
}

// FacultyData model is db model.
type FacultyData struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Code string `db:"code"`
}

func (f *Faculty) ToStorage() *FacultyData {
	return &FacultyData{
		ID:   f.ID,
		Name: f.Name,
		Code: f.Code,
	}
}

func (f *FacultyData) ToServer() *Faculty {
	return &Faculty{
		ID:   f.ID,
		Name: f.Name,
		Code: f.Code,
	}
}

// UpdateInputReq model for updating fields of faculties
type UpdateInputReq struct {
	Name string `form:"faculty-name"`
	Code string `form:"faculty-code"`
}

func (u *UpdateInputReq) ToStorage(facultyID int) *FacultyData {
	return &FacultyData{
		ID:   facultyID,
		Name: u.Name,
		Code: u.Code,
	}
}

func (u UpdateInputReq) Validate() (string, error) {
	if u.Code == "" && u.Name == "" {
		return "update structure has no value", nil
	}

	return "", nil
}

package faculty

// We use in this model type.
// DTO and DAO models.

// DTO model is data transfer object.
type Req struct {
	Name  string `form:"faculty-name" json:"facultyName" validate:"required"`
	Code  string `form:"faculty-code" json:"faculty-code" validate:"required"`
	Index int    `form:"faculty-index" json:"facultyIndex" validate:"required"`
}

// DAO model is data access object.
type Res struct {
	Name  string `db:"name"`
	Code  string `db:"faculty_code"`
	Index int    `db:"faculty_index"`
}

// ToStorage model receives DTO model and save it to the database.
func (d *Req) ToStorage() *Res {
	return &Res{
		Name:  d.Name,
		Code:  d.Code,
		Index: d.Index,
	}
}

// Faculty model is DTO model.
type DTO struct {
	ID    int    `json:"facultyID"`
	Name  string `form:"faculty-name" json:"facultyName" validate:"required"`
	Code  string `form:"faculty-code" json:"faculty-code" validate:"required"`
	Index int    `form:"index-faculty" json:"indexOfFaculty"`
}

// FacultyData model is db model.
type DAO struct {
	ID    int    `db:"id"`
	Name  string `db:"faculty_name"`
	Code  string `db:"faculty_code"`
	Index int    `db:"faculty_index"`
}

func (f *DTO) ToStorage() *DAO {
	return &DAO{
		ID:    f.ID,
		Name:  f.Name,
		Code:  f.Code,
		Index: f.Index,
	}
}

func (f *DAO) ToServer() *DTO {
	return &DTO{
		ID:    f.ID,
		Name:  f.Name,
		Code:  f.Code,
		Index: f.Index,
	}
}

// UpdateInputReq model for updating fields of faculties.
type UpdateInputReq struct {
	Name  string `form:"faculty-name"`
	Code  string `form:"faculty-code"`
	Index int    `form:"faculty-index"`
}

func (u *UpdateInputReq) ToStorage(facultyID int) *DAO {
	return &DAO{
		ID:    facultyID,
		Name:  u.Name,
		Code:  u.Code,
		Index: u.Index,
	}
}

func (u UpdateInputReq) Validate() (string, error) {
	if u.Code == "" && u.Name == "" && u.Index == 0 {
		return "update structure has no value", nil
	}

	return "", nil
}

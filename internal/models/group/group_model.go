package group

// GroupReq model is
type GroupReq struct {
	FacultyID int    `form:"faculty-id" json:"facultyID" validate:"required"`
	ClassCode string `form:"class-code" json:"classCode" validate:"required"`
}

// GroupRes model is
type GroupRes struct {
	FacultyID int    `db:"faculty_id"`
	ClassCode string `db:"class_code"`
}

// ToServer is
func (g *GroupRes) ToServer() *GroupReq {
	return &GroupReq{
		FacultyID: g.FacultyID,
		ClassCode: g.ClassCode,
	}
}

// ToStorage is
func (g *GroupReq) ToStorage() *GroupRes {
	return &GroupRes{
		FacultyID: g.FacultyID,
		ClassCode: g.ClassCode,
	}
}

// GroupData model is
type GroupDAO struct {
	ID        int    `db:"id"`
	FacultyID int    `db:"faculty_id"`
	ClassCode string `db:"class_code"`
}

// Group model is
type GroupDTO struct {
	ID        int    `json:"groupID"`
	FacultyID int    `form:"faculty-id" json:"facultyID"`
	ClassCode string `form:"class-code" json:"classCode"`
}

// ToStorage is
func (g *GroupDTO) ToStorage() *GroupDAO {
	return &GroupDAO{
		ID:        g.ID,
		FacultyID: g.FacultyID,
		ClassCode: g.ClassCode,
	}
}

// ToServer is
func (g *GroupDAO) ToServer() *GroupDTO {
	return &GroupDTO{
		ID:        g.ID,
		FacultyID: g.FacultyID,
		ClassCode: g.ClassCode,
	}
}

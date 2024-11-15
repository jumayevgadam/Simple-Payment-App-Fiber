package group

// GroupReq model is request model for adding group to DB.
type GroupReq struct {
	FacultyID int    `form:"faculty-id" json:"facultyID" validate:"required"`
	ClassCode string `form:"class-code" json:"classCode" validate:"required"`
}

// GroupRes model is response model which get from DB.
type GroupRes struct {
	FacultyID int    `db:"faculty_id"`
	ClassCode string `db:"class_code"`
}

// ToServer method for sending DB model to server.
func (g *GroupRes) ToServer() *GroupReq {
	return &GroupReq{
		FacultyID: g.FacultyID,
		ClassCode: g.ClassCode,
	}
}

// ToStorage method for sending request model to storage.
func (g *GroupReq) ToStorage() *GroupRes {
	return &GroupRes{
		FacultyID: g.FacultyID,
		ClassCode: g.ClassCode,
	}
}

// GroupData model is db model.
type GroupDAO struct {
	ID        int    `db:"id"`
	FacultyID int    `db:"faculty_id"`
	ClassCode string `db:"class_code"`
}

// GroupDTO model for service and handler layer performing request actions.
type GroupDTO struct {
	ID        int    `json:"groupID"`
	FacultyID int    `form:"faculty-id" json:"facultyID"`
	ClassCode string `form:"class-code" json:"classCode"`
}

// ToStorage method for sending DTO model to storage.
func (g *GroupDTO) ToStorage() GroupDAO {
	return GroupDAO{
		ID:        g.ID,
		FacultyID: g.FacultyID,
		ClassCode: g.ClassCode,
	}
}

// ToServer method for sending DAO to server.
func (g *GroupDAO) ToServer() GroupDTO {
	return GroupDTO{
		ID:        g.ID,
		FacultyID: g.FacultyID,
		ClassCode: g.ClassCode,
	}
}

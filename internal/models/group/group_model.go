package group

// GroupReq model is request model for adding group to DB.
type GroupReq struct {
	FacultyID  int    `form:"faculty-id" json:"facultyID" validate:"required"`
	GroupCode  string `form:"group-code" json:"groupCode" validate:"required"`
	CourseYear int    `form:"course-year" json:"courseYear" validate:"required,lte=5,gte=1"`
}

// GroupRes model is response model which get from DB.
type GroupRes struct {
	FacultyID  int    `db:"faculty_id"`
	GroupCode  string `db:"group_code"`
	CourseYear int    `db:"course_year"`
}

// ToServer method for sending DB model to server.
func (g *GroupRes) ToServer() *GroupReq {
	return &GroupReq{
		FacultyID:  g.FacultyID,
		GroupCode:  g.GroupCode,
		CourseYear: g.CourseYear,
	}
}

// ToStorage method for sending request model to storage.
func (g *GroupReq) ToStorage() *GroupRes {
	return &GroupRes{
		FacultyID:  g.FacultyID,
		GroupCode:  g.GroupCode,
		CourseYear: g.CourseYear,
	}
}

// GroupData model is db model.
type GroupDAO struct {
	ID         int    `db:"id"`
	FacultyID  int    `db:"faculty_id"`
	GroupCode  string `db:"group_code"`
	CourseYear int    `db:"course_year"`
}

// GroupDTO model for service and handler layer performing request actions.
type GroupDTO struct {
	ID         int    `json:"groupID"`
	FacultyID  int    `json:"facultyID"`
	GroupCode  string `json:"groupCode"`
	CourseYear int    `json:"courseYear"`
}

// ToStorage method for sending DTO model to storage.
func (g *GroupDTO) ToStorage() *GroupDAO {
	return &GroupDAO{
		ID:         g.ID,
		FacultyID:  g.FacultyID,
		GroupCode:  g.GroupCode,
		CourseYear: g.CourseYear,
	}
}

// ToServer method for sending DAO to server.
func (g *GroupDAO) ToServer() *GroupDTO {
	return &GroupDTO{
		ID:         g.ID,
		FacultyID:  g.FacultyID,
		GroupCode:  g.GroupCode,
		CourseYear: g.CourseYear,
	}
}

// UpdateGroupInput model for updating group model fields.
type UpdateGroupReq struct {
	FacultyID  int    `form:"faculty-id"`
	GroupCode  string `form:"group-code"`
	CourseYear int    `form:"course-year"`
}

func (u *UpdateGroupReq) ToStorage(groupID int) *GroupDAO {
	return &GroupDAO{
		ID:         groupID,
		FacultyID:  u.FacultyID,
		GroupCode:  u.GroupCode,
		CourseYear: u.CourseYear,
	}
}

// If update structure has no value, then must return that.
func (u UpdateGroupReq) Validate() (string, error) {
	if u.GroupCode == "" && u.FacultyID == 0 && u.CourseYear == 0 {
		return "update structure has no value", nil
	}

	return "", nil
}

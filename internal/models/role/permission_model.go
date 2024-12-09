package role

// PermissionReq struct for permissions.
type PermissionReq struct {
	PermissionType string `form:"permission-type" json:"permissionType"`
}

// PermissionRes is data access object.
type PermissionRes struct {
	PermissionType string `db:"permission_type"`
}

// Sent request to storage.
func (p *PermissionReq) ToStorage() PermissionRes {
	return PermissionRes{PermissionType: p.PermissionType}
}

// Sent response to server.
func (p *PermissionRes) ToServer() PermissionReq {
	return PermissionReq{PermissionType: p.PermissionType}
}

// Permission model is dto model.
type Permission struct {
	ID             int    `json:"permission_id"`
	PermissionType string `json:"permission_type"`
}

// PermissionData model is dao model (db model).
type PermissionData struct {
	ID             int    `db:"id"`
	PermissionType string `db:"permission_type"`
}

// Send dto to storage.
func (p *Permission) ToStorage() *PermissionData {
	return &PermissionData{ID: p.ID, PermissionType: p.PermissionType}
}

// Send dao to server.
func (p *PermissionData) ToServer() *Permission {
	return &Permission{ID: p.ID, PermissionType: p.PermissionType}
}

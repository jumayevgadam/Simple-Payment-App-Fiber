package role

// RolePermissionReq model is dto model.
type RolePermissionReq struct {
	RoleID       int `form:"role-id" json:"roleID"`
	PermissionID int `form:"permission-id" json:"permissionID"`
}

// RolePermissionRes model is dao model.
type RolePermissionRes struct {
	RoleID       int `db:"role_id"`
	PermissionID int `db:"permission_id"`
}

// Sending request to storage.
func (r *RolePermissionReq) ToStorage() RolePermissionRes {
	return RolePermissionRes{RoleID: r.RoleID, PermissionID: r.PermissionID}
}

// Sending response to server.
func (r *RolePermissionRes) ToServer() RolePermissionReq {
	return RolePermissionReq{RoleID: r.RoleID, PermissionID: r.PermissionID}
}

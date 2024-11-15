package constants

import "time"

const (
	// nothingUpdatedForRole is.
	NothingUpdatedForRole = "no role name changed"
	// roleUpdateRes is.
	RoleUpdateRes = "role successfully updated"
)

const (
	// RoleTimeDuration need for keeping cached value.
	RoleTimeDuration = 10 * time.Minute
	// ContextTimeOut is.
	ContextTimeout = 10 * time.Second
)

package constants

import "time"

var (
	OnlySuperAdmin = []string{"superadmin"}

	OnlySuperAdminID = []int{1}

	AdminRoles = []string{"superadmin", "admin"}

	AdminRoleIDs = []int{1, 2}

	StudentActionRoles = []string{"superadmin", "student"}

	StudentActionRoleIDs = []int{1, 3}
)

const (
	// when user sign-up to application, it's default roleID is 3.
	DefaultRoleID = 3

	// RoleType.
	RoleType = "role_type"

	// UserID.
	UserID = "user_id"

	// token expiration time.
	TokenExpiryTime = 24 // 1day.

	// SuperAdmin.
	SuperAdmin = "superadmin"

	// Admin.
	Admin = "admin"

	// Student.
	Student = "student"

	// NoUpdateResponse.
	NoUpdateResponse = "update structure has no value"

	// ZeroSevenFiveFive.
	ZeroSevenFiveFive = 0755

	// TransactionTimeOut.
	TransactionTimeOut = 15 * time.Second

	// MinPassword.
	MinPasswordLength = 6

	// ErrMinPasswordLength.
	ErrMinPasswordLength = "updated password must have at least 6 symbol"

	// AtLeastPaymentPrice.
	AtLeastPaymentPrice = 3820

	// FullPaymentPrice.
	FullPaymentPrice = 7920
)

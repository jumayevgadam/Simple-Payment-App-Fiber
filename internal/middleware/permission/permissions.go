package permission

// Payments.
const (
	AddPayment = "add:payment"

	GetPaymentPhoto = "get:payment:photo"

	StudentListPayments = "student:list:payments"
)

// Times.
const (
	AddTime = "add:time"
)

// Faculties.
const (
	CreateFaculty = "create:faculty"

	UpdateFaculty = "update:faculty"

	GetFaculty = "get:faculty"

	ListFaculties = "list:faculties"

	DeleteFaculty = "delete:faculty"

	GetGroupsByFacultyID = "get:groups:by:faculty:id"
)

// Groups.
const (
	CreateGroup = "create:group"

	GetGroup = "get:group"

	ListGroups = "list:groups"

	DeleteGroup = "delete:group"

	UpdateGroup = "update:group"

	ListStudentsByGroupID = "list:students:by:group:id"
)

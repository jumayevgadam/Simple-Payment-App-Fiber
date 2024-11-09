package pagination

// PaginationOps struct is
type PaginationOps struct {
	Limit  int `form:"limit" json:"limit"`
	Offset int `form:"offset" json:"offset"`
}

package abstract

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Pagination struct for sending request pagination ops.
type PaginationQuery struct {
	Limit   int    `form:"limit" json:"limit" validate:"required,gte=0,lt=101"`
	Page    int    `form:"page" json:"page" validate:"required,gte=0"`
	OrderBy string `form:"orderBy" json:"orderBy"`
}

// PaginationData struct is db model, we use in repo layer.
type PaginationData struct {
	Limit   int    `db:"limit"`
	Page    int    `db:"page"`
	OrderBy string `db:"order_by"`
}

// PaginatedResponse model for responsing any model.
type PaginatedResponse[T any] struct {
	Items     []T `json:"items"`
	Limit     int `json:"limit"`
	Page      int `json:"page"`
	TotalPage int `json:"totalPage"`
}

// ToStorage func sends Pagination request to db.
func (p *PaginationQuery) ToStorage() PaginationData {
	return PaginationData{
		Limit:   p.Limit,
		Page:    p.Page,
		OrderBy: p.OrderBy,
	}
}

// we use fiber context for getting params in query Param.
func (pq *PaginationQuery) SetLimit(limit string) error {
	if limit == "" {
		pq.Limit = 10
		return nil
	}

	n, err := strconv.Atoi(limit)
	if err != nil {
		return fmt.Errorf("error in string convert to int: %w", err)
	}
	pq.Limit = n

	return nil
}

func (pq *PaginationQuery) SetPage(page string) error {
	if page == "" {
		pq.Page = 1
		pq.Limit = 10
		return nil
	}

	n, err := strconv.Atoi(page)
	if err != nil {
		return fmt.Errorf("error in string convert to int: %w", err)
	}
	pq.Page = n

	return nil
}

func (pq *PaginationQuery) SetOrderBy(orderBy string) {
	pq.OrderBy = orderBy
}

func GetPaginationFromFiberCtx(c *fiber.Ctx) (PaginationQuery, error) {
	pq := PaginationQuery{}
	if err := pq.SetPage(c.Query("page")); err != nil {
		return pq, fmt.Errorf("error: setting page in query: %w", err)
	}
	if err := pq.SetLimit(c.Query("limit")); err != nil {
		return pq, fmt.Errorf("error: setting limit in query: %w", err)
	}
	pq.SetOrderBy(c.Query("orderBy"))

	return pq, nil
}

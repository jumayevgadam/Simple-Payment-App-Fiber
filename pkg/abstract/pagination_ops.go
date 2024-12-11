package abstract

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// PaginationQuery struct for sending request pagination ops.
type PaginationQuery struct {
	Limit int `form:"limit" json:"limit" validate:"required,gte=0,lt=101"`
	Page  int `form:"page" json:"page" validate:"required,gte=0"`
}

// PaginationData struct is db model, we use in repo layer.
type PaginationData struct {
	Limit int `db:"limit"`
	Page  int `db:"page"`
}

// PaginatedResponse model for uses generics.
type PaginatedResponse[T any] struct {
	Items      []T `json:"items"`
	Limit      int `json:"limit"`
	Page       int `json:"page"`
	TotalItems int `json:"totalItems"`
}

// PaginatedResponseData is db model.
type PaginatedResponseData[T any] struct {
	Items      []T `db:"items"`
	Limit      int `db:"limit"`
	Page       int `db:"page"`
	TotalItems int `db:"total_items"`
}

// ToStorage func sends Pagination request to db.
func (p *PaginationQuery) ToPsqlDBStorage() PaginationData {
	return PaginationData{
		Limit: p.Limit,
		Page:  p.Page,
	}
}

// We use fiber context for getting params in query Param.

func (p *PaginationQuery) SetLimit(limit string) error {
	if limit == "" {
		p.Limit = 10
		return nil
	}

	n, err := strconv.Atoi(limit)
	if err != nil {
		return fmt.Errorf("error in string convert to int: %w", err)
	}
	p.Limit = n

	return nil
}

func (p *PaginationQuery) SetPage(page string) error {
	if page == "" {
		p.Page = 1
		return nil
	}

	n, err := strconv.Atoi(page)
	if err != nil {
		return fmt.Errorf("error in string convert to int: %w", err)
	}
	p.Page = n

	return nil
}

func GetPaginationFromFiberCtx(c *fiber.Ctx) (PaginationQuery, error) {
	pq := PaginationQuery{}
	if err := pq.SetPage(c.Query("page")); err != nil {
		return pq, fmt.Errorf("error: setting page in query: %w", err)
	}

	if err := pq.SetLimit(c.Query("limit")); err != nil {
		return pq, fmt.Errorf("error: setting limit in query: %w", err)
	}

	return pq, nil
}

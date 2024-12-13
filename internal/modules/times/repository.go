package times

import (
	"context"

	timeModel "github.com/jumayevgadam/tsu-toleg/internal/models/time"
)

type Repository interface {
	AddTime(ctx context.Context, response timeModel.Response) (int, error)
}

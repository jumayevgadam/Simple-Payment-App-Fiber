package times

import (
	"context"

	timeModel "github.com/jumayevgadam/tsu-toleg/internal/models/time"
)

type Service interface {
	AddTime(ctx context.Context, request timeModel.Request) (int, error)
}

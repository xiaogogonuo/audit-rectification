package repo

import (
	"audit-rectification/internal/rectserver/model"
	"context"
)

type Repository interface {
	Fetch(ctx context.Context, str string) (model.Response, error)
}

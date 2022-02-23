package service

import (
	"audit-rectification/internal/rectserver/model"
	"context"
)

type Service interface {
	Fetching(ctx context.Context, str string) (model.Response, error)
}

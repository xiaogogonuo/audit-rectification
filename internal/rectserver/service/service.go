package service

import (
	"audit-rectification/internal/rectserver/model"
	"context"
)

type Service interface {
	Fetch(ctx context.Context, str string) (model.Response, error)
}

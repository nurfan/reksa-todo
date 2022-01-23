package repository

import (
	"context"

	"github.com/nurfan/reksa-todo/model"
)

type Psql interface {
	GetUser(ctx context.Context, username string) (result model.Account, err error)
}

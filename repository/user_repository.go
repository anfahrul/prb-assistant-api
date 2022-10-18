package repository

import (
	"context"

	"github.com/anfahrul/prb-assistant-api/entity"
)

type UserRepository interface {
	Insert(ctx context.Context, user entity.User) (entity.User, error)
	LoginCheck(ctx context.Context, user entity.User) (string, error)
	UpdateLoginStatus(ctx context.Context, user entity.User, loginStatus int) error
}

package soc_net

import (
	"context"
	"otus/socNet/internal/structs"
)

type Storage interface {
	GetUser(ctx context.Context, userID int64) (structs.StatusUserDB, error)
	SaveUser(ctx context.Context, userData *structs.User) (int64, error)
	GetHash(ctx context.Context, login *structs.UserLogin) (*structs.UserLogin, error)
	SearchUsers(ctx context.Context, userSearch *structs.UserSearch) ([]*structs.UserSearchRes, error)
}

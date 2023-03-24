package soc_net

import (
	"context"

	"otus/socNet/internal/structs"
)

type SocNetAPI interface {
	GetUser(ctx context.Context, userID int64) (structs.StatusUserDB, error)
	UserRegister(ctx context.Context, user *structs.User) (string, error)
	Login(ctx context.Context, login *structs.UserLogin) (*structs.UserLogin, error)
	UserSearch(ctx context.Context, userSearch *structs.UserSearch) ([]*structs.UserSearchRes, error)
}

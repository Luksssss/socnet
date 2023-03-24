package soc_net

import (
	"context"
	"strconv"

	"otus/socNet/internal/structs"
)

type service struct {
	storage Storage
}

// New DB
func New(storage Storage) *service {
	return &service{
		storage: storage,
	}
}

func (s *service) UserRegister(ctx context.Context, userData *structs.User) (string, error) {
	id, err := s.storage.SaveUser(ctx, userData)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(id, 10), nil
}

func (s *service) GetUser(ctx context.Context, userID int64) (structs.StatusUserDB, error) {
	status, err := s.storage.GetUser(ctx, userID)
	if err != nil {
		return structs.StatusUserDBNotValid, err
	}
	return status, nil
}

func (s *service) Login(ctx context.Context, login *structs.UserLogin) (*structs.UserLogin, error) {
	login, err := s.storage.GetHash(ctx, login)
	if err != nil {
		return nil, err
	}
	return login, nil
}

func (s *service) UserSearch(ctx context.Context, userSearch *structs.UserSearch) ([]*structs.UserSearchRes, error) {
	users, err := s.storage.SearchUsers(ctx, userSearch)
	if err != nil {
		return nil, err
	}
	return users, nil
}

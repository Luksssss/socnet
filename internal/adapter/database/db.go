package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"otus/socNet/database"
	"otus/socNet/internal/structs"
)

type service struct {
	client *database.Storage
}

// New DB
func New(client *database.Storage) *service {
	return &service{
		client: client,
	}
}

func (s *service) SaveUser(ctx context.Context, userData *structs.User) (int64, error) {
	var (
		idUser int64
	)
	query := `insert into users (first_name, second_name, date_birth, id_city, pass)
			  values (
			          $1,
			          $2, 
			          $3, 
			          (select id from city where name = $4),
			          $5
			          ) returning id`

	err := s.client.DB.QueryRow(ctx, query,
		userData.FirstName,
		userData.SecondName,
		userData.DateBirth,
		userData.City,
		userData.Pass,
	).Scan(&idUser)

	if err != nil {
		return 0, err
	}
	return idUser, nil
}

func (s *service) GetUser(ctx context.Context, userID int64) (structs.StatusUserDB, error) {
	var (
		idUser int64
	)

	query := `select id from users where id = $1`
	err := s.client.DB.QueryRow(ctx, query,
		userID,
	).Scan(&idUser)
	if err == pgx.ErrNoRows {
		return structs.StatusUserDBNotFound, nil
	} else if err != nil {
		return structs.StatusUserDBNotValid, err
	}

	return structs.StatusUserDBOK, nil
}
func (s *service) GetHash(ctx context.Context, login *structs.UserLogin) (*structs.UserLogin, error) {
	var (
		hash string
	)

	query := `select pass from users where id = $1`
	err := s.client.DB.QueryRow(ctx, query,
		login.ID,
	).Scan(&hash)
	if err == pgx.ErrNoRows {
		return login, nil
	} else if err != nil {
		return nil, err
	}
	login.Hash = hash
	return login, nil
}

func (s *service) SearchUsers(ctx context.Context, userSearch *structs.UserSearch) ([]*structs.UserSearchRes, error) {
	query := `select first_name, second_name, date_birth from users where 
                                                          lower(first_name) like $1 || '%' AND 
                                                          lower(second_name) like $2 || '%' order by id`
	rows, err := s.client.DB.Query(ctx, query,
		userSearch.FirstName,
		userSearch.SecondName,
	)
	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("невалидные данные")
	} else if err != nil {
		return nil, err
	}
	defer rows.Close()
	resUsers := make([]*structs.UserSearchRes, 0, 2)

	for rows.Next() {
		u := structs.UserSearchRes{}
		err = rows.Scan(&u.FirstName, &u.SecondName, &u.DateBirth)
		if err != nil {
			return nil, err
		}
		resUsers = append(resUsers, &u)
	}
	return resUsers, nil
}

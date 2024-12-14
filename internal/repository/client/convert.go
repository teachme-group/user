package client

import (
	"gitlab.com/coinhubs/balance/internal/domain"
	"gitlab.com/coinhubs/balance/internal/storage/postgres"
)

func userFromRepository(user postgres.User) domain.User {
	return domain.User{
		ID:        user.ID.Bytes,
		Email:     user.Email,
		Login:     user.Login,
		CreatedAt: user.CreatedAt.Time,
	}
}

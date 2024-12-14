package client

import (
	"github.com/teachme-group/user/internal/domain"
	"github.com/teachme-group/user/internal/storage/postgres"
)

func userFromRepository(user postgres.User) domain.User {
	return domain.User{
		ID:        user.ID.Bytes,
		Email:     user.Email,
		Login:     user.Login,
		CreatedAt: user.CreatedAt.Time,
	}
}

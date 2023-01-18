package service

import (
	"context"
	"github.com/Yscream/RESTapi-gin-mongodb/models/user"
)

type RegisterUsers interface {
	GetUser(ctx context.Context, ID uint64) models.User
	FetchUsers()
	InsertUsers()
	UpdateUser()
	DeleteUser()
	DeleteUsers()
}

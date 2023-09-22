package user

import (
	"github.com/skyris/auth-server/internal/usecase/adapters/database"
)

type Options struct{}

type UseCase struct {
	adapterStorage database.User
	options        Options
}

func New(storage database.User, options Options) *UseCase {
	return &UseCase{
		adapterStorage: storage,
		options:        options,
	}
}

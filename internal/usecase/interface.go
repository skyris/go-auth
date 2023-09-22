package usecase

import "github.com/skyris/auth-server/internal/database/dto"

type User interface {
	Register(username, email, password string) (*dto.User, error)
	Login(email, password string) error
}

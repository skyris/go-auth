package database

import (
	"github.com/google/uuid"
	"github.com/skyris/auth-server/internal/database/dto"
)

type User interface {
	Register(username, email, password string) (*dto.User, error)
	GetUserByEmail(email string) (*dto.User, error)
	SoftDeleteUser(id uuid.UUID) error
	HardDeleteUser(id uuid.UUID) error
}

package user

import (
	"github.com/skyris/auth-server/internal/database/dto"
	"log"
)

func (u *UseCase) Register(username, email, password string) (*dto.User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	hashedPasswordString := toBase64(hashedPassword)
	user, err := u.adapterStorage.Register(username, email, hashedPasswordString)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UseCase) Login(email, password string) error {
	user, err := u.adapterStorage.GetUserByEmail(email)
	if err != nil {
		log.Println(err)
		return err
	}
	hashedPassword, err := fromBase64(user.Password)
	if err != nil {
		log.Println(err)
		return err
	}
	err = comparePassword(password, hashedPassword)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

package user

import (
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := "12345678900000000😎"
	hashedPass, err := hashPassword(pass)
	if err != nil {
		fmt.Println(err)
	}
	err = comparePassword(pass, hashedPass)
	if err != nil {
		fmt.Println("not logged in")
		return
	}
	fmt.Println("logged in")
}

func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error while generating bcrypt %w", err)
	}
	return bs, nil
}

func comparePassword(password string, hashedPass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err != nil {
		return fmt.Errorf("invalid password: %w", err)
	}
	return nil
}

func toBase64(message []byte) string {
	return base64.StdEncoding.EncodeToString([]byte(message))
}

func fromBase64(message string) ([]byte, error) {
	return base64.RawStdEncoding.DecodeString(message)
}

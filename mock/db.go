package mock

import "fmt"

// This package is used for to imitate DB queries

type User struct {
	Uid          string `json:"uid"`
	Email        string `json:"email"`
	Password     string `json:"-"`
	PasswordHash string `json:"-"`
}

func GetUser(email, password string) (*User, error) {
	if email == "test" && password == "test" {
		return &User{
			Uid:   "123",
			Email: email,
		}, nil
	}

	return nil, fmt.Errorf("Wrong password")
}

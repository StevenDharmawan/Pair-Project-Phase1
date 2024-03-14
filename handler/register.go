package handler

import (
	"PairProjectPhase1/config"
	"PairProjectPhase1/entity"
	"golang.org/x/crypto/bcrypt"
)

func Register(user entity.User) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = config.DB.Exec("INSERT INTO users(email, password, name, address, phone_number) VALUES (?,?,?,?,?)", user.Email, password, user.Name, user.Address, user.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

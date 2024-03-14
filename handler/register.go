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

	result, err := config.DB.Exec("INSERT INTO users (email, password) VALUES (?, ?)", user.Email, password)
	if err != nil {
		return err
	}
	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	_, err = config.DB.Exec("INSERT INTO user_details (user_id, name, address, phone_number) VALUES (?,?,?,?)", userID, user.Name, user.Address, user.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

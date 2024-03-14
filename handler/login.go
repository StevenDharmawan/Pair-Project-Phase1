package handler

import (
	"PairProjectPhase1/config"
	"PairProjectPhase1/entity"
	"golang.org/x/crypto/bcrypt"
)

func Login(email, password string) (entity.User, bool) {
	var user entity.User
	row := config.DB.QueryRow("SELECT * FROM users WHERE email = ?", email)
	err := row.Scan(&user.UserId, &user.Email, &user.Password, &user.Name, &user.Address, &user.PhoneNumber, &user.Balance, &user.Role)
	if err != nil {
		return user, false
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, false
	}
	return user, true
}

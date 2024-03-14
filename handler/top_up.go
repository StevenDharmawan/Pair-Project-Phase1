package handler

import (
	"PairProjectPhase1/config"
	"PairProjectPhase1/entity"
)

func AddBalance(user entity.User) error {
	_, err := config.DB.Exec("UPDATE users SET balance = balance + ? WHERE user_id = ?", user.Balance, user.UserId)
	if err != nil {
		return err
	}
	return nil
}

package handler

import (
	"PairProjectPhase1/config"
	"PairProjectPhase1/entity"
)

func RemoveCart(user entity.User, userInput int) error {
	_, err := config.DB.Exec("DELETE FROM carts WHERE user_id IN (SELECT user_id FROM (SELECT user_id FROM carts WHERE user_id = ? ORDER BY product_detail_id LIMIT 1 OFFSET ?) AS subquery)", user.UserId, userInput)
	if err != nil {
		return err
	}
	return nil
}

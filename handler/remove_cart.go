package handler

import (
	"PairProjectPhase1/config"
	"PairProjectPhase1/entity"
)

func RemoveCart(user entity.User, product entity.Product) error {
	_, err := config.DB.Exec("DELETE FROM carts WHERE user_id = ? AND product_detail_id = ? ORDER BY product_detail_id", user.UserId, product.ProductDetailId)
	if err != nil {
		return err
	}
	return nil
}

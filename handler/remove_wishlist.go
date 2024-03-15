package handler

import (
	"PairProjectPhase1/config"
	"PairProjectPhase1/entity"
)

func RemoveWishlist(user entity.User, product entity.Product) error {
	_, err := config.DB.Exec("DELETE FROM wishlists WHERE user_id = ? AND product_detail_id = ? ORDER BY product_detail_id", user.UserId, product.ProductDetailId)
	if err != nil {
		return err
	}
	return nil
}

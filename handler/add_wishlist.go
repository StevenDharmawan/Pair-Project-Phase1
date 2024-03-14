package handler

import (
	"PairProjectPhase1/config"
	"PairProjectPhase1/entity"
)

func AddWishlist(user entity.User, product entity.Products) error {
	_, err := config.DB.Exec("INSERT INTO wishlists(user_id, product_detail_id) VALUES (?,?)", user.UserId, product.ProductDetailId)
	if err != nil {
		return err
	}
	return nil
}

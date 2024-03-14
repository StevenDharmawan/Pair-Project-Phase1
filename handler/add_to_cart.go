package handler

import (
	"PairProjectPhase1/config"
	"PairProjectPhase1/entity"
)

func AddToCart(user entity.User, product entity.Products, cart entity.Cart) error {
	row := config.DB.QueryRow("SELECT price FROM products_detail WHERE product_detail_id = ?", product.ProductDetailId)
	err := row.Scan(&product.Price)
	if err != nil {
		return err
	}
	_, err = config.DB.Exec("INSERT INTO carts(user_id, product_detail_id, quantity, total) VALUES (?,?,?,?)", user.UserId, product.ProductDetailId, cart.Quantity, float64(cart.Quantity)*product.Price)
	if err != nil {
		return err
	}
	return nil
}

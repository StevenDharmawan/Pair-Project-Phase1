package handler

import (
	"PairProjectPhase1/config"
	"PairProjectPhase1/entity"
	"errors"
)

func AddToCart(user entity.User, product entity.Product, cart entity.Cart) error {
	row := config.DB.QueryRow("SELECT price, stock FROM products_detail WHERE product_detail_id = ?", product.ProductDetailId)
	err := row.Scan(&product.Price, &product.Stock)
	if err != nil {
		return err
	}
	if product.Stock < cart.Quantity {
		return errors.New("quantity melebihi stock yang ada")
	}
	_, err = config.DB.Exec("INSERT INTO carts(user_id, product_detail_id, quantity, total) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE quantity = quantity + ?, total = total + ?", user.UserId, product.ProductDetailId, cart.Quantity, float64(cart.Quantity)*product.Price, cart.Quantity, float64(cart.Quantity)*product.Price)
	if err != nil {
		return err
	}
	return nil
}

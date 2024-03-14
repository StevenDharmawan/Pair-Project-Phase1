package handler

import (
	"PairProjectPhase1/config"
	"PairProjectPhase1/entity"
)

func UpdateStock(product entity.Products) error {
	_, err := config.DB.Exec("UPDATE products_detail SET stock = stock + ? WHERE product_detail_id = ?", product.Stock, product.ProductDetailId)
	if err != nil {
		return err
	}
	return nil
}

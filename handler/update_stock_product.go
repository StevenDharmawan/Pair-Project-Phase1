package handler

import (
	"PairProjectPhase1/config"
)

func UpdateStock(productID int, newStock int) error {
	_, err := config.DB.Exec("UPDATE products_detail SET stock = ? WHERE product_id = ?", newStock, productID)
	if err != nil {
		return err
	}
	config.DB.Close()
	return nil
}

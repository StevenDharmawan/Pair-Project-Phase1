package handler

import (
	"PairProjectPhase1/config"
	"PairProjectPhase1/entity"
)

func AddProduct(product entity.Product) error {
	result, err := config.DB.Exec("INSERT IGNORE INTO products(product_name, description) VALUES (?,?)", product.ProductName, product.Description)
	if err != nil {
		return err
	}
	productID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	if productID == 0 {
		row := config.DB.QueryRow("SELECT product_id FROM products WHERE product_name = ?", product.ProductName)
		err := row.Scan(&product.ProductId)
		if err != nil {
			return err
		}
		_, err = config.DB.Exec("INSERT INTO products_detail(product_id, size_id, price, stock) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE stock = stock + ?", product.ProductId, product.SizeId, product.Price, product.Stock)
		if err != nil {
			return err
		}
	} else {
		_, err = config.DB.Exec("INSERT INTO products_detail(product_id, size_id, price, stock) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE stock = stock + ?", productID, product.SizeId, product.Price, product.Stock)
		if err != nil {
			return err
		}
	}

	return nil
}

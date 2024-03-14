package handler

import (
	"PairProjectPhase1/config"
	"database/sql"
	"fmt"
)

func StockReport() {
	query := `SELECT p.product_name, s.size_name, pd.stock
	          FROM products_detail pd
	          JOIN products p ON pd.product_id = p.product_id
	          JOIN sizes s ON pd.size_id = s.size_id
	          ORDER BY pd.stock ASC` //agar bisa urut jumlah tekecil

	rows, err := config.DB.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err.Error())
		}
	}(rows)

	// Print the table headers
	fmt.Println("Product Stock")
	fmt.Printf("| %-20s | %-10s | %-10s |\n", "Product Name", "Size", "Stock")
	fmt.Println("|----------------------|------------|------------|")

	for rows.Next() {
		var productName, sizeName string
		var stock int
		err := rows.Scan(&productName, &sizeName, &stock)
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("| %-20s | %-10s | %-10d |\n", productName, sizeName, stock)
	}
}

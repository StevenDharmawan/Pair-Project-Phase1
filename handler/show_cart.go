package handler

import (
	"PairProjectPhase1/config"
	"PairProjectPhase1/entity"
	"fmt"
)

func ShowCart(user entity.User) {
	query := `
        SELECT c.product_detail_id, p.product_name, s.size_name, c.quantity, c.total
        FROM carts c 
        JOIN products_detail pd ON c.product_detail_id = pd.product_detail_id
        JOIN products p ON pd.product_id = p.product_id 
        JOIN sizes s ON pd.size_id = s.size_id 
		WHERE user_id = ?
        ORDER BY c.product_detail_id;
    `
	rows, err := config.DB.Query(query, user.UserId)
	if err != nil {
		panic(err.Error())
	}

	// Header tabel
	fmt.Printf("| %-5s | %-30s | %-10s | %-14s | %-5s |\n", "No", "Product Name", "Size Name", "Quantity", "Total")
	fmt.Println("|-------|--------------------------------|------------|----------------|-------|")

	for rows.Next() {
		var name, sizeName string
		var productDetailId, quantity int
		var total float64
		err = rows.Scan(&productDetailId, &name, &sizeName, &quantity, &total)
		if err != nil {
			panic(err.Error())
		}

		// Cetak baris tabel
		fmt.Printf("| %-5d | %-30s | %-10s | Rp.%-11.2f | %-5d |\n", productDetailId, name, sizeName, total, quantity)
	}
}

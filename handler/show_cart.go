package handler

import (
	"PairProjectPhase1/config"
	"fmt"
)

func ShowCart() {
	query := `
        SELECT p.product_name, s.size_name, c.quantity, c.total
        FROM carts c 
        JOIN products_detail pd ON c.product_detail_id = pd.product_detail_id
        JOIN products p ON pd.product_id = p.product_id 
        JOIN sizes s ON pd.size_id = s.size_id;
    `
	rows, err := config.DB.Query(query)
	if err != nil {
		panic(err.Error())
	}

	// Header tabel
	fmt.Printf("| %-5s | %-30s | %-10s | %-14s | %-5s |\n", "No", "Product Name", "Size Name", "Quantity", "Total")
	fmt.Println("|-------|--------------------------------|------------|----------------|-------|")

	// Variabel untuk nomor urut
	var count int

	for rows.Next() {
		var name, sizeName string
		var quantity int
		var total float64
		err = rows.Scan(&name, &sizeName, &quantity, &total)
		if err != nil {
			panic(err.Error())
		}

		// Increment nomor urut
		count++

		// Cetak baris tabel
		fmt.Printf("| %-5d | %-30s | %-10s | Rp.%-11.2f | %-5d |\n", count, name, sizeName, total, quantity)
	}
}

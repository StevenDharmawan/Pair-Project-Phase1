package handler

import (
	"PairProjectPhase1/config"
	"fmt"
)

func HistoryOrder() {
	query := `
		SELECT p.product_name, s.size_name, c.quantity, c.total, oh.OrderDate
		FROM carts c
		JOIN order_history oh ON c.cart_id = oh.cart_id
		JOIN products_detail pd ON c.product_detail_id = pd.product_detail_id
		JOIN products p ON pd.product_id = p.product_id
		JOIN sizes s ON pd.size_id = s.size_id;
	`
	rows, err := config.DB.Query(query)
	if err != nil {
		panic(err.Error())
	}

	// Header tabel
	fmt.Printf("| %-5s | %-30s | %-10s | %-8s | %-6s | %-10s |\n", "No", "Product Name", "Size Name", "Quantity", "Total", "Order Date")
	fmt.Println("|-------|--------------------------------|------------|----------|--------|------------|")

	// Variabel untuk nomor urut
	var count int

	for rows.Next() {
		var name, sizeName string
		var quantity, total int
		var OrderDate string // Gunakan tipe string untuk tanggal
		err = rows.Scan(&name, &sizeName, &quantity, &total, &OrderDate)
		if err != nil {
			panic(err.Error())
		}

		// Increment nomor urut
		count++

		// Cetak baris tabel
		fmt.Printf("| %-5d | %-30s | %-10s | %-8d | Rp.%-4.2f | %-10s |\n", count, name, sizeName, quantity, total, OrderDate)
	}
}

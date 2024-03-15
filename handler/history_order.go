package handler

import (
	"PairProjectPhase1/config"
	"PairProjectPhase1/entity"
	"fmt"
)

func HistoryOrder(user entity.User) {
	query := "SELECT product_detail_id, quantity, total, order_date FROM order_history WHERE user_id = ?"
	rows, err := config.DB.Query(query, user.UserId)
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

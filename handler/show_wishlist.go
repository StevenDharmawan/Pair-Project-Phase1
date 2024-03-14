package handler

import (
	"PairProjectPhase1/config"
	"fmt"
)

func ShowWishlist() {
	query := `
        SELECT p.product_name, s.size_name
        FROM wishlists w
        JOIN products_detail pd ON w.product_detail_id = pd.product_detail_id
        JOIN products p ON pd.product_id = p.product_id 
        JOIN sizes s ON pd.size_id = s.size_id;
    `
	rows, err := config.DB.Query(query)
	if err != nil {
		panic(err.Error())
	}

	// Header tabel
	fmt.Printf("| %-5s | %-30s | %-10s |\n", "No", "Product Name", "Size Name")
	fmt.Println("|-------|--------------------------------|------------|")

	// Variabel untuk nomor urut
	var count int

	for rows.Next() {
		var name, sizeName string
		err = rows.Scan(&name, &sizeName)
		if err != nil {
			panic(err.Error())
		}

		// Increment nomor urut
		count++

		// Cetak baris tabel
		fmt.Printf("| %-5d | %-30s | %-10s |\n", count, name, sizeName)
	}
}

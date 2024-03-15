package handler

import (
	"PairProjectPhase1/config"
	"PairProjectPhase1/entity"
	"fmt"
)

func ShowWishlist(user entity.User) {
	query := `
        SELECT pd.product_detail_id, p.product_name, s.size_name
        FROM wishlists w
        JOIN products_detail pd ON w.product_detail_id = pd.product_detail_id
        JOIN products p ON pd.product_id = p.product_id 
        JOIN sizes s ON pd.size_id = s.size_id
		WHERE user_id = ?;
    `
	rows, err := config.DB.Query(query, user.UserId)
	if err != nil {
		panic(err.Error())
	}

	// Header tabel
	fmt.Printf("| %-5s | %-30s | %-10s |\n", "No", "Product Name", "Size Name")
	fmt.Println("|-------|--------------------------------|------------|")

	for rows.Next() {
		var productDetailId int
		var name, sizeName string
		err = rows.Scan(&productDetailId, &name, &sizeName)
		if err != nil {
			panic(err.Error())
		}
		// Cetak baris tabel
		fmt.Printf("| %-5d | %-30s | %-10s |\n", productDetailId, name, sizeName)
	}
	fmt.Println()
}

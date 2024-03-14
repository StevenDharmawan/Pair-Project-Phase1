package handler

import (
	"PairProjectPhase1/config"
	"database/sql"
	"fmt"
)

func OrderReport() {
	query := `SELECT oh.order_history_id, oh.cart_id, oh.OrderDate, COUNT(*) AS total_sales
	          FROM order_history oh
	          JOIN carts c ON oh.cart_id = c.cart_id
	          GROUP BY oh.cart_id`
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
	fmt.Println("Order History")
	fmt.Printf("| %-5s | %-10s | %-20s | %-15s |\n", "No", "Cart ID", "Order Date", "Total Sales")
	fmt.Println("|-------|------------|----------------------|-----------------|")

	var count int

	for rows.Next() {
		var orderID, cartID, totalSales int
		var orderDate string
		err := rows.Scan(&orderID, &cartID, &orderDate, &totalSales)
		if err != nil {
			panic(err.Error())
		}

		count++

		fmt.Printf("| %-5d | %-10d | %-20s | %-15d |\n", count, cartID, orderDate, totalSales)
	}
}

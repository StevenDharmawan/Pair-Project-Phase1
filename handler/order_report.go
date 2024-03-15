package handler

import (
	"PairProjectPhase1/config"
	"database/sql"
	"fmt"
)

func OrderReport() {
	query := `SELECT user_details.name, SUM(total) AS total_sales 
				FROM order_history
				JOIN user_details ON order_history.user_id = user_details.user_id
				GROUP BY user_id`
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
	fmt.Printf("| %-5s | %-15s | %-10s |\n", "No", "Name", "Total")
	fmt.Println("|-------|-------------------|------------|")

	var count int

	for rows.Next() {
		var total float64
		var name string
		err := rows.Scan(&name, &total)
		if err != nil {
			panic(err.Error())
		}

		count++

		fmt.Printf("| %-5s | %-15s | %-10s |\n", count, name, total)
	}
}

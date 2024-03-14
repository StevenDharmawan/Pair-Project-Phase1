package handler

import (
	"PairProjectPhase1/config"
	"fmt"
)

func ReportTotal() {
	query := "SELECT name, email FROM users WHERE user_id != 1"

	// Execute the query
	rows, err := config.DB.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Print the table headers
	fmt.Println("Total Users")
	fmt.Printf("| %-5s | %-20s | %-70s |\n", "No", "Name", "Email")
	fmt.Println("|-------|----------------------|------------------------------------------------------------------------|")

	var count int

	for rows.Next() {
		var name, email string
		err := rows.Scan(&name, &email)
		if err != nil {
			panic(err.Error())
		}

		count++ // Increment the row count

		// Print the row data in a formatted manner
		fmt.Printf("| %-5d | %-20s | %-70s |\n", count, name, email)
	}
	fmt.Println()
}

package handler

import (
    "PairProjectPhase1/config"
    "fmt"
)

func ShowProduct() {
    query := `
        SELECT p.product_name, p.description, s.size_name, pd.price, pd.stock
        FROM products p 
        JOIN products_detail pd ON p.product_id = pd.product_id 
        JOIN sizes s ON pd.size_id = s.size_id;
    `
    rows, err := config.DB.Query(query)
    if err != nil {
        panic(err.Error())
    }

    // Header tabel
    fmt.Printf("| %-5s | %-20s | %-70s | %-10s | %-14s | %-5s |\n", "No", "Name", "Description", "Size Name", "Price", "Stock")
    fmt.Println("|-------|----------------------|------------------------------------------------------------------------|------------|----------------|-------|")

    // Variabel untuk nomor urut
    var count int

    for rows.Next() {
        var name, description, sizeName string
        var stock int
        var price float64
        err = rows.Scan(&name, &description, &sizeName, &price, &stock)
        if err != nil {
            panic(err.Error())
        }

        // Increment nomor urut
        count++

        // Cetak baris tabel
        fmt.Printf("| %-5d | %-20s | %-70s | %-10s | Rp.%-11.2f | %-5d |\n", count, name, description, sizeName, price, stock)
    }
}
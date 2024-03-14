import (
	"PairProjectPhase1/config"
	"fmt"
)

func ShowProduct() {
	query := `SELECT p.product_name, p.description, s.size_name, pd.price, pd.stock FROM products p JOIN products_detail pd ON p.product_id = pd.product_id JOIN sizes s ON pd.size_id = s.size_id;`
	rows, err := config.DB.Query(query)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var name, description, size_name string
		var stock int
		var price float64
		err = rows.Scan(&name, &description, &size_name, &price, &stock)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("Name: ", name)
		fmt.Println("Description: ", description)
		fmt.Println("Size Name: ", size_name)
		fmt.Println("Price: ", price)
		fmt.Println("Stock: ", stock)
		fmt.Println("---------------------------------------------")

	}
}

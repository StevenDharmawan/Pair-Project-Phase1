package handler

import (
	"PairProjectPhase1/config"
	"PairProjectPhase1/entity"
	"errors"
	"fmt"
	"time"
)

func Checkout(user entity.User, product entity.Product, cart entity.Cart, order entity.Order) error {
	row := config.DB.QueryRow("SELECT SUM(total) FROM carts WHERE user_id = ? GROUP BY user_id", user.UserId)
	err := row.Scan(&cart.Total)
	//err := row.Scan(&user.UserId, &product.ProductDetailId, &cart.Quantity, &user.Name, &user.Address, &user.PhoneNumber, &user.Balance, &user.Role)
	if err != nil {
		return err
	}
	if cart.Total > user.Balance {
		return errors.New("saldo anda tidak cukup")
	} else {
		_, err := config.DB.Exec("UPDATE user_details SET balance = balance - ?", cart.Total)
		if err != nil {
			return err
		}
		query := "SELECT user_id, product_detail_id, quantity, total FROM carts"
		rows, err := config.DB.Query(query, user.UserId)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("asd")
		for rows.Next() {
			err = rows.Scan(&user.UserId, &product.ProductDetailId, &cart.Quantity, &cart.Total)
			if err != nil {
				panic(err.Error())
			}
			order.OrderTime = time.Now()
			_, err = config.DB.Exec("INSERT INTO order_history (user_id, product_detail_id, quantity, total, order_date) VALUES (?,?,?,?,?)", user.UserId, product.ProductDetailId, cart.Quantity, cart.Total, order.OrderTime)
			if err != nil {
				return err
			}
		}
		_, err = config.DB.Exec("DELETE FROM carts WHERE user_id = ?", user.UserId)
		if err != nil {
			return err
		}
	}
	return nil
}

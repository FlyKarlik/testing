package repository

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	db *sqlx.DB
}

func NewDb(db *sqlx.DB) *DB {
	return &DB{db: db}
}
func (d *DB) GetOrders(orders []string) error {
	query := `
    SELECT p.name, p.id, o.order_id, po.quantity, s.name, ps.is_main, ps.additional_shelves
    FROM products p, products_orders po, orders o, products_shelves ps, shelves s
    WHERE p.id = po.id_product
    AND po.id_order = o.order_id
    AND p.id = ps.id_product
    AND ps.id_shelf = s.id
    AND o.order_id IN `
	orderIDs := "("
	for i, arg := range orders {
		orderID, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalf("Неверный номер заказа: %s", arg)
		}
		orderIDs += strconv.Itoa(orderID)
		if i < len(orders)-1 {
			orderIDs += ","
		}
	}
	orderIDs += ")"

	rows, err := d.db.Query(query + orderIDs)
	if err != nil {
		return err
	}
	defer rows.Close()

	shelfMap := make(map[string][]string)

	for rows.Next() {
		var productName, shelfName, additionalShelves string
		var productID, orderID, quantity int
		var isMain bool
		err := rows.Scan(&productName, &productID, &orderID, &quantity, &shelfName, &isMain, &additionalShelves)
		if err != nil {
			return err
		}

		productInfo := fmt.Sprintf("%s (id=%d)\nзаказ %d, %d шт\n", productName, productID, orderID, quantity)
		if !isMain {
			productInfo += fmt.Sprintf("доп стеллаж: %s\n", additionalShelves)
		}
		shelfMap[shelfName] = append(shelfMap[shelfName], productInfo)
	}
	fmt.Printf("Страница сборки заказов %s\n", orderIDs)
	fmt.Println()
	for shelf, products := range shelfMap {
		fmt.Printf("===%s\n", shelf)
		for _, product := range products {
			fmt.Println(product)
		}
	}

	return nil
}

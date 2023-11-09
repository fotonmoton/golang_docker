package db

import (
	"database/sql"
	"docker/warehouse"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlWarehouseState struct {
	db *sql.DB
}

func NewMysqlConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/warehouse")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func NewMysqlWarehouseState() *MysqlWarehouseState {
	return &MysqlWarehouseState{
		NewMysqlConnection(),
	}
}

func (w *MysqlWarehouseState) List() []warehouse.Product {
	products := []warehouse.Product{}
	rows, err := w.db.Query("SELECT * FROM products")

	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var p warehouse.Product
		if err := rows.Scan(&p.Id, &p.Name, &p.Qty); err != nil {
			log.Println(err)

			return nil
		}
		products = append(products, p)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)

		return nil
	}

	return products
}

func (w *MysqlWarehouseState) Save(p warehouse.Product) int {
	result, err := w.db.Exec("INSERT INTO products (name, qty) VALUES (?, ?)", p.Name, p.Qty)
	if err != nil {
		log.Println(err)
		return 0
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0
	}
	return int(id)
}
